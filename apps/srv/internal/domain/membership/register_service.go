package membership

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type RegisterService struct {
	userStore UserStore
	Hasher    Hasher
}

func NewRegisterService(userStore UserStore, hasher Hasher) *RegisterService {
	return &RegisterService{
		userStore: userStore,
		Hasher:    hasher,
	}
}

func (s *RegisterService) Register(ctx context.Context, user User) (User, error) {
	existingUser, err := s.userStore.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return User{}, fmt.Errorf("failed to check if user exists by email: %w", err)
	}

	// the user exists...
	if existingUser.ID != uuid.Nil {
		return existingUser, nil
	}

	hashedPassword, err := s.Hasher.Hash(ctx, user.Password)
	if err != nil {
		return User{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = hashedPassword
	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = time.Now().UTC()
	createdUser, err := s.userStore.CreateUser(ctx, user)
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return createdUser, nil
}
