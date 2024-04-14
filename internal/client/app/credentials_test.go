//go:build integration
// +build integration

package app_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
	"github.com/Galish/goph-keeper/internal/entity"

	"github.com/stretchr/testify/assert"
)

var errVersionRequired = errors.New("version is required")

func TestCredentials(t *testing.T) {
	var id string
	username := generateUsername()
	password := generatePassword()

	client := grpc.NewClient(cfg, auth.New())

	userUseCase := user.New(client)
	notesUseCase := notes.New(client)

	t.Run("should fail with authorization error", func(t *testing.T) {
		list, err := notesUseCase.GetCredentialsList(context.Background())

		assert.Nil(t, list)
		assert.Equal(t, notes.ErrAuthRequired, err)
	})

	t.Run("should successfully authenticate the user", func(t *testing.T) {
		err := userUseCase.SignUp(context.Background(), username, password)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentialsList(context.Background())

		assert.Empty(t, creds)
		assert.NoError(t, err)
	})

	t.Run("should fail with validation error", func(t *testing.T) {
		creds := &entity.Credentials{
			Title:       "Gmail",
			Description: fmt.Sprintf("Gmail account credentials for %s", username),
			Username:    username,
		}

		err := notesUseCase.AddCredentials(context.Background(), creds)

		assert.Equal(t, notes.ErrInvalidEntity, err)
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

	t.Run("should return note by id", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentials(context.Background(), id)

		expected := &entity.Credentials{
			Title:       "Gmail",
			Description: fmt.Sprintf("Gmail account credentials for %s", username),
			Username:    username,
			Password:    password,
			Version:     0,
		}

		assert.Equal(t, expected, creds)
		assert.NoError(t, err)
	})

	t.Run("should fail with validation error", func(t *testing.T) {
		creds := &entity.Credentials{
			ID:       id,
			Title:    "Yandex",
			Username: username,
		}

		err := notesUseCase.UpdateCredentials(context.Background(), creds, false)

		assert.Equal(t, notes.ErrInvalidEntity, err)
	})

	t.Run("should fail with version conflict error", func(t *testing.T) {
		creds := &entity.Credentials{
			ID:          id,
			Title:       "Yandex",
			Description: "Yandex mail account credentials",
			Username:    username,
			Password:    password,
		}

		err := notesUseCase.UpdateCredentials(context.Background(), creds, false)

		assert.Equal(t, errVersionRequired, err)
	})

	t.Run("should successfully update note", func(t *testing.T) {
		creds := &entity.Credentials{
			ID:          id,
			Title:       "Yandex",
			Description: "Yandex mail account credentials",
			Username:    username,
			Password:    password,
			Version:     1,
		}

		err := notesUseCase.UpdateCredentials(context.Background(), creds, false)

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

		err := notesUseCase.UpdateCredentials(context.Background(), creds, true)

		assert.NoError(t, err)
	})

	t.Run("should return undated note by identifier", func(t *testing.T) {
		creds, err := notesUseCase.GetCredentials(context.Background(), id)

		expected := &entity.Credentials{
			Title:       "Yandex",
			Description: fmt.Sprintf("Yandex mail account credentials for %s", username),
			Username:    username,
			Password:    password,
			Version:     2,
		}

		assert.Equal(t, expected, creds)
		assert.NoError(t, err)
	})

	t.Run("should successfully delete the note", func(t *testing.T) {
		err := notesUseCase.DeleteCredentials(context.Background(), id)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		list, err := notesUseCase.GetCredentialsList(context.Background())

		assert.Empty(t, list)
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
