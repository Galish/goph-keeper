package cli_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
	"github.com/Galish/goph-keeper/internal/entity"

	"github.com/stretchr/testify/assert"
)

var (
	id       string
	username string = randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 10)
	password string = randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890.-!", 25)
)

func randString(symbols string, n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}

	return string(b)
}

func TestAddCredentials(t *testing.T) {
	cfg := config.New()

	grpcClient := grpc.NewClient(cfg, auth.New())

	userUseCase := user.New(grpcClient)
	notesUseCase := notes.New(grpcClient)

	t.Run("should fail with authorization error", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentialsList(context.Background())

		assert.Nil(t, creds)
		assert.Equal(t, notes.ErrAuthRequired, err)
	})

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

	t.Run("should return an empty list", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentialsList(context.Background())

		assert.Empty(t, creds)
		assert.NoError(t, err)
	})

	t.Run("should successfully create a note", func(t *testing.T) {
		creds := &entity.Credentials{
			Title:       "Gmail",
			Description: fmt.Sprintf("Gmail account credentials for %s", username),
			Username:    username,
			Password:    password,
		}

		err := notesUseCase.AddCredentials(context.Background(), creds)

		assert.NoError(t, err)
	})

	t.Run("should return a list containing the newly added note", func(t *testing.T) {
		list, err := notesUseCase.GetCredentialsList(context.Background())
		id = list[0].ID

		assert.Equal(t, 1, len(list))
		assert.Equal(t, "Gmail", list[0].Title)
		assert.Equal(t, fmt.Sprintf("Gmail account credentials for %s", username), list[0].Description)
		assert.NoError(t, err)
	})

	t.Run("should return a note by identifier", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentials(context.Background(), id)

		assert.Equal(t, &entity.Credentials{
			Title:       "Gmail",
			Description: fmt.Sprintf("Gmail account credentials for %s", username),
			Username:    username,
			Password:    password,
			Version:     0,
		}, creds)

		assert.NoError(t, err)
	})

	t.Run("should successfully update note", func(t *testing.T) {
		creds := &entity.Credentials{
			ID:          id,
			Title:       "Yandex",
			Description: fmt.Sprintf("Yandex mail account credentials for %s", username),
			Username:    username,
			Password:    password,
		}

		err := notesUseCase.UpdateCredentials(context.Background(), creds, true) // overwrite

		assert.NoError(t, err)
	})

	t.Run("should return undated note by identifier", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentials(context.Background(), id)

		assert.Equal(t, &entity.Credentials{
			Title:       "Yandex",
			Description: fmt.Sprintf("Yandex mail account credentials for %s", username),
			Username:    username,
			Password:    password,
			Version:     1,
		}, creds)

		assert.NoError(t, err)
	})

	t.Run("should successfully delete the note", func(t *testing.T) {
		err := notesUseCase.DeleteCredentials(context.Background(), id)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentialsList(context.Background())

		assert.Empty(t, creds)
		assert.NoError(t, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentials(context.Background(), id)

		assert.Nil(t, creds)
		assert.Equal(t, notes.ErrNotFound, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		creds := &entity.Credentials{
			ID:          id,
			Title:       "Gmail",
			Description: fmt.Sprintf("Gmail account credentials for %s", username),
			Username:    username,
			Password:    password,
		}

		err := notesUseCase.UpdateCredentials(context.Background(), creds, true)

		assert.Equal(t, notes.ErrNotFound, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		err := notesUseCase.DeleteCredentials(context.Background(), id)

		assert.Equal(t, notes.ErrNotFound, err)
	})
}
