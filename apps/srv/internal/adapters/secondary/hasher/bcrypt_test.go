//go:build unit

package hasher_test

import (
	"context"
	"strings"
	"testing"

	"github.com/arthurdotwork/bastion/internal/adapters/secondary/hasher"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestBcryptHasher_Hash(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h := hasher.NewBcryptHasher(bcrypt.DefaultCost)

	t.Run("it should return an error if it can not hash the password", func(t *testing.T) {
		t.Parallel()

		_, err := h.Hash(ctx, strings.Repeat("a", 128))
		require.Error(t, err)
	})

	t.Run("it should hash the password", func(t *testing.T) {
		t.Parallel()

		hash, err := h.Hash(ctx, "password")
		require.NoError(t, err)
		require.NotEmpty(t, hash)
		require.NotEqual(t, "password", hash)
	})
}

func TestBcryptHasher_Verify(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	h := hasher.NewBcryptHasher(bcrypt.DefaultCost)

	t.Run("it should return an error if the password does not match", func(t *testing.T) {
		hash, err := h.Hash(ctx, "password")
		require.NoError(t, err)

		err = h.Verify(ctx, "wrong-password", hash)
		require.Error(t, err)
	})

	t.Run("it should verify the password", func(t *testing.T) {
		hash, err := h.Hash(ctx, "password")
		require.NoError(t, err)

		err = h.Verify(ctx, "password", hash)
		require.NoError(t, err)
	})
}
