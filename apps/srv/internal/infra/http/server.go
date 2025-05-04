package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	*gin.Engine

	addr           string
	instantiatedAt time.Time
}

func NewServer(addr string, allowedOrigins []string) *Server {
	srv := gin.New()
	srv.Use(useCors(allowedOrigins))
	srv.Use(gin.Recovery())

	return &Server{
		Engine:         srv,
		addr:           addr,
		instantiatedAt: time.Now().UTC(),
	}
}

func (s *Server) Serve(ctx context.Context) error {
	s.GET("/checks/liveness", s.livenessProbe())
	s.GET("/checks/readiness", s.readinessProbe(ctx))

	srv := &http.Server{
		Addr:              s.addr,
		Handler:           s,
		ReadTimeout:       time.Second * 2,
		WriteTimeout:      time.Second * 2,
		ReadHeaderTimeout: time.Second * 2,
	}

	grp, _ := errgroup.WithContext(ctx)
	grp.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return srv.Shutdown(ctx)
	})

	grp.Go(func() error {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		return nil
	})

	return grp.Wait()
}

func uptime(since time.Duration) string {
	return fmt.Sprintf("%.2fs", since.Seconds())
}

func (s *Server) livenessProbe() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "uptime": uptime(time.Since(s.instantiatedAt))})
	}
}

func (s *Server) readinessProbe(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		select {
		case <-ctx.Done():
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "unhealthy"})
			return
		default:
			c.JSON(http.StatusOK, gin.H{"status": "healthy", "uptime": uptime(time.Since(s.instantiatedAt))})
		}
	}
}
