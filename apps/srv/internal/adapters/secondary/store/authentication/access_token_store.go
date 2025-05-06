package authentication

import (
	"context"

	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
)

type AccessTokenStore struct {
	db psql.Queryable
	q  *queries.Queries
}

func NewAccessTokenStore(db psql.Queryable, q *queries.Queries) *AccessTokenStore {
	return &AccessTokenStore{
		db: db,
		q:  q,
	}
}

func (s *AccessTokenStore) CreateAccessToken(ctx context.Context, token authentication.AccessToken) (authentication.AccessToken, error) {
	createAccessTokenParams := queries.CreateAccessTokenParams{
		ID:              token.ID,
		UserID:          token.UserID,
		TokenIdentifier: token.TokenIdentifier,
		IssuedAt:        token.IssuedAt,
		ExpiresAt:       token.ExpiresAt,
	}

	createdToken, err := s.q.CreateAccessToken(ctx, createAccessTokenParams)
	if err != nil {
		return authentication.AccessToken{}, err
	}

	return authentication.AccessToken{
		ID:              createdToken.ID,
		UserID:          createdToken.UserID,
		TokenIdentifier: createdToken.TokenIdentifier,
		IssuedAt:        createdToken.IssuedAt,
		ExpiresAt:       createdToken.ExpiresAt,
		MaxAge:          token.MaxAge,
		RawToken:        token.RawToken,
	}, nil
}
