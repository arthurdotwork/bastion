package container

import (
	"context"
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/arthurdotwork/bastion/internal/adapters/primary/http/handler"
	"github.com/arthurdotwork/bastion/internal/adapters/secondary/hasher"
	"github.com/arthurdotwork/bastion/internal/adapters/secondary/paseto"
	authenticationStore "github.com/arthurdotwork/bastion/internal/adapters/secondary/store/authentication"
	membershipStore "github.com/arthurdotwork/bastion/internal/adapters/secondary/store/membership"
	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/arthurdotwork/bastion/internal/domain/membership"
	"github.com/arthurdotwork/bastion/internal/infra/http"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Container struct {
	ctx context.Context

	dependenciesMap     *sync.Map
	shutdownFuncs       []func() error
	initializationError chan error
}

func New(ctx context.Context) *Container {
	return &Container{
		ctx:                 ctx,
		dependenciesMap:     &sync.Map{},
		shutdownFuncs:       []func() error{},
		initializationError: make(chan error, 1),
	}
}

func (c *Container) SetupDatabase() *psql.DB {
	return singleton(c, func() *psql.DB {
		db, err := psql.Connect(
			c.ctx,
			env("DATABASE_USERNAME", "postgres"),
			env("DATABASE_PASSWORD", "postgres"),
			env("DATABASE_HOST", "localhost"),
			env("DATABASE_PORT", "5432"),
			env("DATABASE_NAME", "postgres"),
		)
		if err != nil {
			c.recordInitializationError(fmt.Errorf("could not connect to database: %w", err))
			return nil
		}

		c.shutdownFuncs = append(c.shutdownFuncs, db.Close)
		return db
	})
}

func (c *Container) SetupQueries() *queries.Queries {
	return singleton(c, func() *queries.Queries {
		q, err := queries.Prepare(c.ctx, c.SetupDatabase())
		if err != nil {
			c.recordInitializationError(fmt.Errorf("could not prepare queries: %w", err))
			return nil
		}

		c.shutdownFuncs = append(c.shutdownFuncs, q.Close)
		return q
	})
}

func (c *Container) SetupHTTPServer() *http.Server {
	return singleton(c, func() *http.Server {
		// TODO: Replace allowed origins with a proper configuration.
		srv := http.NewServer(env("HTTP_ADDR", ":8080"), []string{"http://localhost:5173"})
		return srv
	})
}

func (c *Container) SetupRegisterHandler() gin.HandlerFunc {
	return singleton(c, func() gin.HandlerFunc {
		return handler.Register(c.SetupMembershipRegisterService())
	})
}

func (c *Container) SetupAuthenticationHandler() gin.HandlerFunc {
	return singleton(c, func() gin.HandlerFunc {
		return handler.Authenticate(c.SetupAuthenticationService())
	})
}

func (c *Container) SetupVerifyAuthenticationHandler() gin.HandlerFunc {
	return singleton(c, func() gin.HandlerFunc {
		return handler.VerifyAuthentication(c.SetupAuthenticationService())
	})
}

func (c *Container) SetupAuthenticationService() *authentication.Service {
	return singleton(c, func() *authentication.Service {
		return authentication.NewService(c.SetupAuthenticationUserStore(), c.SetupBcryptHasher(), c.SetupPasetoProvider(), c.SetupAuthenticationAccessTokenStore())
	})
}

func (c *Container) SetupAuthenticationUserStore() *authenticationStore.UserStore {
	return singleton(c, func() *authenticationStore.UserStore {
		return authenticationStore.NewUserStore(c.SetupDatabase(), c.SetupQueries())
	})
}

func (c *Container) SetupAuthenticationAccessTokenStore() *authenticationStore.AccessTokenStore {
	return singleton(c, func() *authenticationStore.AccessTokenStore {
		return authenticationStore.NewAccessTokenStore(c.SetupDatabase(), c.SetupQueries())
	})
}

func (c *Container) SetupPasetoProvider() *paseto.Provider {
	return singleton(c, func() *paseto.Provider {
		pasetoSecretKey, err := base64.StdEncoding.DecodeString(env("PASETO_SECRET_KEY", ""))
		if err != nil {
			c.recordInitializationError(fmt.Errorf("could not decode PASETO secret key: %w", err))
			return nil
		}

		return paseto.NewProvider(pasetoSecretKey)
	})
}

func (c *Container) SetupUserStore() *membershipStore.UserStore {
	return singleton(c, func() *membershipStore.UserStore {
		return membershipStore.NewUserStore(c.SetupDatabase(), c.SetupQueries())
	})
}

func (c *Container) SetupBcryptHasher() *hasher.BcryptHasher {
	return singleton(c, func() *hasher.BcryptHasher {
		return hasher.NewBcryptHasher(bcrypt.DefaultCost)
	})
}

func (c *Container) SetupMembershipRegisterService() *membership.RegisterService {
	return singleton(c, func() *membership.RegisterService {
		return membership.NewRegisterService(c.SetupUserStore(), c.SetupBcryptHasher())
	})
}

func (c *Container) Shutdown() {
	for _, shutdownFunc := range c.shutdownFuncs {
		if err := shutdownFunc(); err != nil {
			slog.ErrorContext(c.ctx, "error during shutdown", "error", err)
		}
	}
}

func (c *Container) recordInitializationError(err error) {
	c.initializationError <- err
}

func (c *Container) InitializationErrorChannel() <-chan error {
	return c.initializationError
}

func singleton[T any](c *Container, factory func() T) T {
	methodName := getCallerMethodName()

	if value, found := c.dependenciesMap.Load(methodName); found {
		return value.(T)
	}

	instance := factory()
	c.dependenciesMap.Store(methodName, instance)

	return instance
}

func getCallerMethodName() string {
	pc, _, _, _ := runtime.Caller(2)
	function := runtime.FuncForPC(pc).Name()
	parts := strings.Split(function, ".")

	return parts[len(parts)-1]
}

func env(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
