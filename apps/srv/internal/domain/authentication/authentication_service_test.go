//go:build unit

package authentication_test

import (
	"context"
	"testing"

	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/arthurdotwork/bastion/internal/domain/authentication/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_AuthenticateWithPassword(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userStore := mocks.NewUserStore(t)
	hasher := mocks.NewHasher(t)
	tokenProvider := mocks.NewTokenProvider(t)
	accessTokenStore := mocks.NewAccessTokenStore(t)

	authenticationService := authentication.NewService(userStore, hasher, tokenProvider, accessTokenStore)

	t.Run("it should return an error if it can not find the user", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(authentication.User{}, assert.AnError).Once()

		_, err := authenticationService.AuthenticateWithPassword(ctx, "email@bastion.dev", "password")
		require.Error(t, err)
	})

	t.Run("it should return an error if the password is incorrect", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(authentication.User{HashedPassword: "password"}, nil).Once()
		hasher.EXPECT().Verify(ctx, "password", "password").Return(assert.AnError).Once()

		_, err := authenticationService.AuthenticateWithPassword(ctx, "email@bastion.dev", "password")
		require.Error(t, err)
	})

	t.Run("it should return an error if the token provider fails to create a token", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(authentication.User{HashedPassword: "password"}, nil).Once()
		hasher.EXPECT().Verify(ctx, "password", "password").Return(nil).Once()
		tokenProvider.EXPECT().Generate(ctx, authentication.User{HashedPassword: "password"}).Return(authentication.AccessToken{}, assert.AnError).Once()

		_, err := authenticationService.AuthenticateWithPassword(ctx, "email@bastion.dev", "password")
		require.Error(t, err)
	})

	t.Run("it should return an error if the access token store fails to save the token", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(authentication.User{HashedPassword: "password"}, nil).Once()
		hasher.EXPECT().Verify(ctx, "password", "password").Return(nil).Once()
		tokenProvider.EXPECT().Generate(ctx, authentication.User{HashedPassword: "password"}).Return(authentication.AccessToken{}, nil).Once()
		accessTokenStore.EXPECT().CreateAccessToken(ctx, authentication.AccessToken{}).Return(authentication.AccessToken{}, assert.AnError).Once()

		_, err := authenticationService.AuthenticateWithPassword(ctx, "email@bastion.dev", "password")
		require.Error(t, err)
	})

	t.Run("it should authenticate the user", func(t *testing.T) {
		userStore.EXPECT().GetUserByEmail(ctx, "email@bastion.dev").Return(authentication.User{HashedPassword: "password"}, nil).Once()
		hasher.EXPECT().Verify(ctx, "password", "password").Return(nil).Once()
		tokenProvider.EXPECT().Generate(ctx, authentication.User{HashedPassword: "password"}).Return(authentication.AccessToken{}, nil).Once()
		accessTokenStore.EXPECT().CreateAccessToken(ctx, authentication.AccessToken{}).Return(authentication.AccessToken{RawToken: "rawToken"}, nil).Once()

		accessToken, err := authenticationService.AuthenticateWithPassword(ctx, "email@bastion.dev", "password")
		require.NoError(t, err)
		require.NotEmpty(t, accessToken)
		require.Equal(t, "rawToken", accessToken.RawToken)
	})
}

func TestService_VerifyAccessToken(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userStore := mocks.NewUserStore(t)
	hasher := mocks.NewHasher(t)
	tokenProvider := mocks.NewTokenProvider(t)
	accessTokenStore := mocks.NewAccessTokenStore(t)

	authenticationService := authentication.NewService(userStore, hasher, tokenProvider, accessTokenStore)

	t.Run("it should return an error if the token is invalid", func(t *testing.T) {
		tokenProvider.EXPECT().Verify(ctx, "invalidToken").Return(assert.AnError).Once()

		err := authenticationService.VerifyAccessToken(ctx, "invalidToken")
		require.Error(t, err)
	})

	t.Run("it should verify the token", func(t *testing.T) {
		tokenProvider.EXPECT().Verify(ctx, "validToken").Return(nil).Once()

		err := authenticationService.VerifyAccessToken(ctx, "validToken")
		require.NoError(t, err)
	})
}
