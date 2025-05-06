package authentication

import (
	"context"

	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/arthurdotwork/bastion/internal/infra/psql"
	"github.com/arthurdotwork/bastion/internal/infra/queries"
)

type UserStore struct {
	db psql.Queryable
	q  *queries.Queries
}

func NewUserStore(db psql.Queryable, q *queries.Queries) *UserStore {
	return &UserStore{
		db: db,
		q:  q,
	}
}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (authentication.User, error) {
	getUserByEmailParams := queries.GetUserByEmailParams{
		Email: email,
	}

	user, err := s.q.GetUserByEmail(ctx, getUserByEmailParams)
	if err != nil {
		return authentication.User{}, err
	}

	return authentication.User{
		ID:             user.ID,
		Email:          user.Email,
		HashedPassword: user.Password,
	}, nil
}
