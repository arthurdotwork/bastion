package authentication

import "context"

type UserStore interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type AccessTokenStore interface {
	CreateAccessToken(ctx context.Context, token AccessToken) (AccessToken, error)
}

type TokenProvider interface {
	Generate(ctx context.Context, user User) (AccessToken, error)
	Verify(ctx context.Context, token string) error
}

type TokenVerifier interface{}

type Hasher interface {
	Verify(ctx context.Context, password, hash string) error
}
