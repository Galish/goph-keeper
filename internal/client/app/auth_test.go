//go:build integration
// +build integration

package app_test

import (
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"

	"github.com/stretchr/testify/assert"
)

var cfg = &config.Config{
	GRPCServAddr: ":3200",
}

func TestAuthentication(t *testing.T) {
	username := generateUsername()
	password := generatePassword()

	client := grpc.NewClient(cfg, auth.New())

	userUseCase := user.New(client)

	t.Run("should fail with user not found error", func(t *testing.T) {
		err := userUseCase.SignIn(context.Background(), username, password)

		assert.Equal(t, user.ErrNotFound, err)
	})

	t.Run("should successfully authenticate the user", func(t *testing.T) {
		err := userUseCase.SignUp(context.Background(), username, password)

		assert.NoError(t, err)
	})

	t.Run("should fail with conflict error", func(t *testing.T) {
		err := userUseCase.SignUp(context.Background(), username, password)

		assert.Equal(t, user.ErrAlreadyExists, err)
	})

	t.Run("should successfully authorize the user", func(t *testing.T) {
		err := userUseCase.SignIn(context.Background(), username, password)

		assert.NoError(t, err)
	})
}

func generateUsername() string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 10)
}

func generatePassword() string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.-!", 25)
}

func randString(symbols string, n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)

	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(b)
}
