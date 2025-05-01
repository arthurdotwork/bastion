package membership

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
