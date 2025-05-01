package hasher

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) *BcryptHasher {
	return &BcryptHasher{
		cost: cost,
	}
}

func (h *BcryptHasher) Hash(ctx context.Context, password string) (string, error) {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", err
	}

	return string(bcryptHash), nil
}
