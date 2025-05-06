package authentication

import (
	"context"
	"fmt"
)

type Service struct {
	userStore        UserStore
	hasher           Hasher
	tokenProvider    TokenProvider
	accessTokenStore AccessTokenStore
}

func NewService(
	userStore UserStore,
	hasher Hasher,
	tokenProvider TokenProvider,
	accessTokenStore AccessTokenStore,
) *Service {
	return &Service{
		userStore:        userStore,
		hasher:           hasher,
		tokenProvider:    tokenProvider,
		accessTokenStore: accessTokenStore,
	}
}

func (s *Service) AuthenticateWithPassword(ctx context.Context, email string, password string) (AccessToken, error) {
	user, err := s.userStore.GetUserByEmail(ctx, email)
	if err != nil {
		return AccessToken{}, fmt.Errorf("user not found: %w", err)
	}

	if err := s.hasher.Verify(ctx, password, user.HashedPassword); err != nil {
		return AccessToken{}, fmt.Errorf("invalid password: %w", err)
	}

	accessToken, err := s.tokenProvider.Generate(ctx, user)
	if err != nil {
		return AccessToken{}, fmt.Errorf("failed to generate access token: %w", err)
	}

	storedAccessToken, err := s.accessTokenStore.CreateAccessToken(ctx, accessToken)
	if err != nil {
		return AccessToken{}, fmt.Errorf("failed to store access token: %w", err)
	}

	return storedAccessToken, nil
}

func (s *Service) VerifyAccessToken(ctx context.Context, rawAccessToken string) error {
	if err := s.tokenProvider.Verify(ctx, rawAccessToken); err != nil {
		return fmt.Errorf("invalid access token: %w", err)
	}

	return nil
}
