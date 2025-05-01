//go:build integration

package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/arthurdotwork/bastion/internal/adapters/secondary/store"
	"github.com/arthurdotwork/bastion/internal/domain/membership"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
	"github.com/stretchr/testify/require"
)

func TestUserStore_CreateUser(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := psql.Connect(ctx, "postgres", "postgres", "localhost", "5432", "postgres")
	require.NoError(t, err)

	tx, err := db.BeginTxx(ctx, nil)
	require.NoError(t, err)
	defer tx.Rollback() //nolint:errcheck

	q := queries.New(tx.Tx())
	userStore := store.NewUserStore(tx, q)

	t.Run("it should create the user", func(t *testing.T) {
		user := membership.User{
			Email:     "email@bastion.dev",
			Password:  "password",
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		}

		createdUser, err := userStore.CreateUser(ctx, user)
		require.NoError(t, err)
		require.NotEmpty(t, createdUser.ID)
		require.EqualValues(t, user.Email, createdUser.Email)
		require.EqualValues(t, user.Password, createdUser.Password)
		require.WithinDuration(t, user.CreatedAt, createdUser.CreatedAt, time.Second)
		require.WithinDuration(t, user.UpdatedAt, createdUser.UpdatedAt, time.Second)
		require.Nil(t, createdUser.DeletedAt)

		t.Run("it should return an error if the user already exists", func(t *testing.T) {
			_, err := userStore.CreateUser(ctx, user)
			require.Error(t, err)
		})
	})
}

func TestUserStore_GetUserByEmail(t *testing.T) {

}
