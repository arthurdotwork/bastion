package paseto

import (
	"context"
	"fmt"
	"time"

	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/google/uuid"
	"github.com/o1egl/paseto/v2"
)

const (
	defaultAccessTokenLifetime = 2 * time.Minute
	audience                   = "bastion"
	issuer                     = "bastion"
)

type Provider struct {
	secretKey []byte
}

func NewProvider(secretKey []byte) *Provider {
	return &Provider{
		secretKey: secretKey,
	}
}

func (p *Provider) Generate(ctx context.Context, user authentication.User) (authentication.AccessToken, error) {
	now := time.Now().UTC()
	exp := now.Add(defaultAccessTokenLifetime)
	nbt := now.Add(-time.Minute)

	tokenIdentifier, err := uuid.NewV7()
	if err != nil {
		return authentication.AccessToken{}, fmt.Errorf("failed to generate token identifier: %w", err)
	}

	pasetoToken := paseto.JSONToken{
		Audience:   audience,
		Issuer:     issuer,
		Jti:        tokenIdentifier.String(),
		Subject:    fmt.Sprintf("user:%s", user.ID),
		Expiration: exp,
		IssuedAt:   now,
		NotBefore:  nbt,
	}
	pasetoToken.Set("purpose", "access")

	pasetoRawToken, err := paseto.NewV2().Encrypt(p.secretKey, pasetoToken, nil)
	if err != nil {
		return authentication.AccessToken{}, fmt.Errorf("failed to encrypt token: %w", err)
	}

	return authentication.AccessToken{
		ID:              uuid.New(),
		UserID:          user.ID,
		TokenIdentifier: tokenIdentifier,
		IssuedAt:        now,
		ExpiresAt:       exp,
		MaxAge:          defaultAccessTokenLifetime.Milliseconds() / 1000,
		RawToken:        pasetoRawToken,
	}, nil
}

func (p *Provider) Verify(ctx context.Context, rawToken string) error {
	var token paseto.JSONToken
	if err := paseto.NewV2().Decrypt(rawToken, p.secretKey, &token, nil); err != nil {
		return fmt.Errorf("failed to decrypt token: %w", err)
	}

	if token.Expiration.Before(time.Now().UTC()) {
		return fmt.Errorf("token is expired")
	}

	return nil
}
