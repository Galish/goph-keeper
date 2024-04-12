//go:build integration
// +build integration

package app_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
	"github.com/Galish/goph-keeper/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestRawNotes(t *testing.T) {
	var id string
	username := generateUsername()
	password := generatePassword()

	client := grpc.NewClient(cfg, auth.New())

	userUseCase := user.New(client)
	notesUseCase := notes.New(client)

	t.Run("should fail with authorization error", func(t *testing.T) {
		list, err := notesUseCase.GetRawNotesList(context.Background())

		assert.Nil(t, list)
		assert.Equal(t, notes.ErrAuthRequired, err)
	})

	t.Run("should successfully authenticate the user", func(t *testing.T) {
		err := userUseCase.SignUp(context.Background(), username, password)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		list, err := notesUseCase.GetRawNotesList(context.Background())

		assert.Empty(t, list)
		assert.NoError(t, err)
	})

	t.Run("should fail with validation error", func(t *testing.T) {
		note := &entity.RawNote{
			Title:       "My secret file",
			Description: "Super secret file data",
		}

		err := notesUseCase.AddRawNote(context.Background(), note)

		assert.Equal(t, notes.ErrInvalidEntity, err)
	})

	t.Run("should successfully create a note", func(t *testing.T) {
		note := &entity.RawNote{
			Title:       "My secret file",
			Description: "Super secret file data",
			Value:       []byte("Super secret file data"),
		}

		err := notesUseCase.AddRawNote(context.Background(), note)

		assert.NoError(t, err)
	})

	t.Run("should return a list containing the newly added note", func(t *testing.T) {
		list, err := notesUseCase.GetRawNotesList(context.Background())
		id = list[0].ID

		assert.Equal(t, 1, len(list))
		assert.Equal(t, "My secret file", list[0].Title)
		assert.Equal(t, "Super secret file data", list[0].Description)
		assert.NoError(t, err)
	})

	t.Run("should return note by id", func(t *testing.T) {
		note, err := notesUseCase.GetRawNote(context.Background(), id)

		expected := &entity.RawNote{
			Title:       "My secret file",
			Description: "Super secret file data",
			Value:       []byte("Super secret file data"),
			Version:     0,
		}

		assert.Equal(t, expected, note)
		assert.NoError(t, err)
	})

	t.Run("should fail with validation error", func(t *testing.T) {
		note := &entity.RawNote{
			ID:          id,
			Title:       "My secret file",
			Description: "Super secret file data",
		}

		err := notesUseCase.UpdateRawNote(context.Background(), note, false)

		assert.Equal(t, notes.ErrInvalidEntity, err)
	})

	t.Run("should fail with version conflict error", func(t *testing.T) {
		note := &entity.RawNote{
			ID:          id,
			Title:       "My secret file",
			Description: "",
			Value:       []byte(fmt.Sprintf("Super secret file data uploaded by %s", username)),
		}

		err := notesUseCase.UpdateRawNote(context.Background(), note, false)

		assert.Equal(t, errVersionRequired, err)
	})

	t.Run("should successfully update note", func(t *testing.T) {
		note := &entity.RawNote{
			ID:          id,
			Title:       "My secret file",
			Description: "",
			Value:       []byte(fmt.Sprintf("Super secret file data uploaded by %s", username)),
			Version:     1,
		}

		err := notesUseCase.UpdateRawNote(context.Background(), note, false)

		assert.NoError(t, err)
	})

	t.Run("should successfully update note", func(t *testing.T) {
		note := &entity.RawNote{
			ID:          id,
			Title:       "My secret file",
			Description: "",
			Value:       []byte(fmt.Sprintf("Super secret file data uploaded by %s", username)),
		}

		err := notesUseCase.UpdateRawNote(context.Background(), note, true)

		assert.NoError(t, err)
	})

	t.Run("should return undated note by identifier", func(t *testing.T) {
		note, err := notesUseCase.GetRawNote(context.Background(), id)

		expected := &entity.RawNote{
			Title:       "My secret file",
			Description: "",
			Value:       []byte(fmt.Sprintf("Super secret file data uploaded by %s", username)),
			Version:     2,
		}

		assert.Equal(t, expected, note)
		assert.NoError(t, err)
	})

	t.Run("should successfully delete the note", func(t *testing.T) {
		err := notesUseCase.DeleteRawNote(context.Background(), id)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		list, err := notesUseCase.GetRawNotesList(context.Background())

		assert.Empty(t, list)
		assert.NoError(t, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		note, err := notesUseCase.GetRawNote(context.Background(), id)

		assert.Nil(t, note)
		assert.Equal(t, notes.ErrNotFound, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		note := &entity.RawNote{
			ID:          id,
			Title:       "My secret file",
			Description: "",
			Value:       []byte("Super secret file data"),
		}

		err := notesUseCase.UpdateRawNote(context.Background(), note, true)

		assert.Equal(t, notes.ErrNotFound, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		err := notesUseCase.DeleteRawNote(context.Background(), id)

		assert.Equal(t, notes.ErrNotFound, err)
	})
}
