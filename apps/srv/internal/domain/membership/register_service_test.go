//go:build unit

package membership_test

import (
	"context"
	"testing"

	"github.com/arthurdotwork/bastion/internal/domain/membership"
	"github.com/arthurdotwork/bastion/internal/domain/membership/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterService_Register(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userStore := mocks.NewUserStore(t)
	hasher := mocks.NewHasher(t)
	registerService := membership.NewRegisterService(userStore, hasher)

	t.Run("it should return an error if it can not retrieve the user", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(membership.User{}, assert.AnError).Once()

		_, err := registerService.Register(ctx, membership.User{Email: "email@bastion.dev"})
		require.Error(t, err)
	})

	t.Run("it should return the existing user if it exists", func(t *testing.T) {
		existingUser := membership.User{ID: uuid.New()}
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(existingUser, nil).Once()

		user, err := registerService.Register(ctx, membership.User{Email: "email@bastion.dev"})
		require.NoError(t, err)
		require.Equal(t, existingUser.ID, user.ID)
	})

	t.Run("it should return an error if it can not hash the password", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(membership.User{}, nil).Once()
		hasher.EXPECT().Hash(ctx, "password").Return("", assert.AnError).Once()

		_, err := registerService.Register(ctx, membership.User{Email: "email@bastion.dev", Password: "password"})
		require.Error(t, err)
	})

	t.Run("it should return an error if it can not create the user", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(membership.User{}, nil).Once()
		hasher.EXPECT().Hash(ctx, "password").Return("hashedPassword", nil).Once()
		userStore.EXPECT().CreateUser(ctx, mock.Anything).Return(membership.User{}, assert.AnError).Once()

		_, err := registerService.Register(ctx, membership.User{Email: "email@bastion.dev", Password: "password"})
		require.Error(t, err)
	})

	t.Run("it should create the user", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(membership.User{}, nil).Once()
		hasher.EXPECT().Hash(ctx, "password").Return("hashedPassword", nil).Once()
		userStore.EXPECT().CreateUser(ctx, mock.Anything).Return(membership.User{ID: uuid.New()}, nil).Once()

		user, err := registerService.Register(ctx, membership.User{Email: "email@bastion.dev", Password: "password"})
		require.NoError(t, err)
		require.NotEmpty(t, user.ID)
	})
}
