package authentication

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
}

type AccessToken struct {
	ID              uuid.UUID
	UserID          uuid.UUID
	TokenIdentifier uuid.UUID
	IssuedAt        time.Time
	ExpiresAt       time.Time
	RevokedAt       *time.Time

	MaxAge   int64
	RawToken string
}
