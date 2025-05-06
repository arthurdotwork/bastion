//go:build integration

package authentication_test

import (
	"context"
	"testing"
	"time"

	authenticationStore "github.com/arthurdotwork/bastion/internal/adapters/secondary/store/authentication"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
	"github.com/stretchr/testify/require"
)

func TestUserStore_GetUserByEmail(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := psql.Connect(ctx, "postgres", "postgres", "localhost", "5432", "postgres")
	require.NoError(t, err)

	tx, err := db.BeginTxx(ctx, nil)
	require.NoError(t, err)
	defer tx.Rollback() //nolint:errcheck

	q := queries.New(tx.Tx())
	accessTokenStore := authenticationStore.NewUserStore(tx, q)

	t.Run("it should return an error if the user can not be found", func(t *testing.T) {
		_, err := accessTokenStore.GetUserByEmail(ctx, "email@bastion.dev")
		require.Error(t, err)
	})

	t.Run("it should return the user", func(t *testing.T) {
		user, err := q.CreateUser(ctx, queries.CreateUserParams{
			Email:     "email@bastion.dev",
			Password:  "password",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		})
		require.NoError(t, err)

		foundUser, err := accessTokenStore.GetUserByEmail(ctx, user.Email)
		require.NoError(t, err)
		require.EqualValues(t, user.ID, foundUser.ID)
	})
}
