//go:build unit

package paseto_test

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/arthurdotwork/bastion/internal/adapters/secondary/paseto"
	"github.com/arthurdotwork/bastion/internal/domain/authentication"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestProvider_Generate(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pasetoSecretKey, err := base64.StdEncoding.DecodeString("LqpJ8gcjveMsVEBzmiLq/w8rhZ8V/t4ceW3cToyK+j0=")
	require.NoError(t, err)

	provider := paseto.NewProvider(pasetoSecretKey)

	t.Run("it should generate a valid access token", func(t *testing.T) {
		user := authentication.User{ID: uuid.New()}

		pasetoToken, err := provider.Generate(ctx, user)
		require.NoError(t, err)
		require.NotEmpty(t, pasetoToken.RawToken)
	})
}

func TestProvider_Verify(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pasetoSecretKey, err := base64.StdEncoding.DecodeString("LqpJ8gcjveMsVEBzmiLq/w8rhZ8V/t4ceW3cToyK+j0=")
	require.NoError(t, err)

	provider := paseto.NewProvider(pasetoSecretKey)

	user := authentication.User{ID: uuid.New()}
	pasetoToken, err := provider.Generate(ctx, user)
	require.NoError(t, err)

	t.Run("it should validate a valid access token", func(t *testing.T) {
		err := provider.Verify(ctx, pasetoToken.RawToken)
		require.NoError(t, err)
	})
}
