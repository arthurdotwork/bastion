package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/arthurdotwork/bastion/internal/domain/membership"
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

func (s *UserStore) Atomic(ctx context.Context, fn func(ctx context.Context, userStore membership.UserStore) error) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	q := s.q.WithTx(tx.Tx().Tx)
	userStore := NewUserStore(tx, q)
	if err := fn(ctx, userStore); err != nil {
		_ = tx.Rollback() //nolint:errcheck
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *UserStore) CreateUser(ctx context.Context, user membership.User) (membership.User, error) {
	createUserParams := queries.CreateUserParams{
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	createdUser, err := s.q.CreateUser(ctx, createUserParams)
	if err != nil {
		return membership.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return membership.User{
		ID:        createdUser.ID,
		Email:     createdUser.Email,
		Password:  createdUser.Password,
		Username:  createdUser.Username,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		DeletedAt: nil,
	}, nil
}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (membership.User, error) {
	getUserByEmailParams := queries.GetUserByEmailParams{Email: email}

	user, err := s.q.GetUserByEmail(ctx, getUserByEmailParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return membership.User{}, nil
		}

		return membership.User{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	var deletedAt *time.Time
	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	return membership.User{
		ID:        user.ID,
		Email:     user.Email,
		Password:  user.Password,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}, nil
}
