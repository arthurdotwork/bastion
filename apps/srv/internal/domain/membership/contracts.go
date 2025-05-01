package membership

import "context"

type UserStore interface {
	Atomic(ctx context.Context, fn func(ctx context.Context, userStore UserStore) error) error
	CreateUser(ctx context.Context, user User) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type Hasher interface {
	Hash(ctx context.Context, password string) (string, error)
}
