//go:build integration

package authentication_test

import (
	"context"
	"testing"
	"time"

	authenticationStore "github.com/arthurdotwork/bastion/internal/adapters/secondary/store/authentication"
	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAccessTokenStore_CreateAccessToken(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := psql.Connect(ctx, "postgres", "postgres", "localhost", "5432", "postgres")
	require.NoError(t, err)

	tx, err := db.BeginTxx(ctx, nil)
	require.NoError(t, err)
	defer tx.Rollback() //nolint:errcheck

	q := queries.New(tx.Tx())
	accessTokenStore := authenticationStore.NewAccessTokenStore(tx, q)

	t.Run("it should create the token", func(t *testing.T) {
		user, err := q.CreateUser(ctx, queries.CreateUserParams{
			Email:     "email@bastion.dev",
			Password:  "password",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})
		require.NoError(t, err)

		token := authentication.AccessToken{
			ID:              uuid.New(),
			UserID:          user.ID,
			TokenIdentifier: uuid.New(),
			IssuedAt:        time.Now().UTC(),
			ExpiresAt:       time.Now().UTC(),
			RawToken:        "token",
		}

		createdToken, err := accessTokenStore.CreateAccessToken(ctx, token)
		require.NoError(t, err)
		require.NotEmpty(t, createdToken.ID)
	})

	t.Run("it should return an error if it can not create the token", func(t *testing.T) {
		token := authentication.AccessToken{
			ID:              uuid.New(),
			UserID:          uuid.UUID{},
			TokenIdentifier: uuid.New(),
			IssuedAt:        time.Now().UTC(),
			ExpiresAt:       time.Now().UTC(),
			RawToken:        "token",
		}

		_, err := accessTokenStore.CreateAccessToken(ctx, token)
		require.Error(t, err)
	})
}
