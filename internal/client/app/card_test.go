//go:build integration
// +build integration

package app_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
	"github.com/Galish/goph-keeper/internal/entity"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	var id string
	username := generateUsername()
	password := generatePassword()

	client := grpc.NewClient(cfg, auth.New())

	userUseCase := user.New(client)
	notesUseCase := notes.New(client)

	t.Run("should fail with authorization error", func(t *testing.T) {
		list, err := notesUseCase.GetCardsList(context.Background())

		assert.Nil(t, list)
		assert.Equal(t, notes.ErrAuthRequired, err)
	})

	t.Run("should successfully authenticate the user", func(t *testing.T) {
		err := userUseCase.SignUp(context.Background(), username, password)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		card, err := notesUseCase.GetCardsList(context.Background())

		assert.Empty(t, card)
		assert.NoError(t, err)
	})

	t.Run("should fail with validation error", func(t *testing.T) {
		card := &entity.Card{
			Title:       "Sberbank",
			Description: "Credit card",
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		}

		err := notesUseCase.AddCard(context.Background(), card)

		assert.Equal(t, notes.ErrInvalidEntity, err)
	})

	t.Run("should successfully create a note", func(t *testing.T) {
		card := &entity.Card{
			Title:       "Sberbank",
			Description: "Credit card",
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		}

		err := notesUseCase.AddCard(context.Background(), card)

		assert.NoError(t, err)
	})

	t.Run("should return a list containing the newly added note", func(t *testing.T) {
		list, err := notesUseCase.GetCardsList(context.Background())
		id = list[0].ID

		assert.Equal(t, 1, len(list))
		assert.Equal(t, "Sberbank", list[0].Title)
		assert.Equal(t, "Credit card", list[0].Description)
		assert.NoError(t, err)
	})

	t.Run("should return note by id", func(t *testing.T) {
		card, err := notesUseCase.GetCard(context.Background(), id)

		expected := &entity.Card{
			Title:       "Sberbank",
			Description: "Credit card",
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			Version:     0,
		}

		assert.Equal(t, expected, card)
		assert.NoError(t, err)
	})

	t.Run("should fail with validation error", func(t *testing.T) {
		card := &entity.Card{
			ID:          id,
			Title:       "Sberbank",
			Description: fmt.Sprintf("Credit card of %s", username),
			Number:      "1234 5678 9012 4453",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		}

		err := notesUseCase.UpdateCard(context.Background(), card, false)

		assert.Equal(t, notes.ErrInvalidEntity, err)
	})

	t.Run("should fail with version conflict error", func(t *testing.T) {
		card := &entity.Card{
			ID:          id,
			Title:       "Sberbank",
			Description: fmt.Sprintf("Credit card of %s", username),
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		}

		err := notesUseCase.UpdateCard(context.Background(), card, false)

		assert.Equal(t, errVersionRequired, err)
	})

	t.Run("should successfully update note", func(t *testing.T) {
		card := &entity.Card{
			ID:          id,
			Title:       "Sberbank",
			Description: fmt.Sprintf("Credit card of %s", username),
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			Version:     1,
		}

		err := notesUseCase.UpdateCard(context.Background(), card, false)

		assert.NoError(t, err)
	})

	t.Run("should successfully update note", func(t *testing.T) {
		card := &entity.Card{
			ID:          id,
			Title:       "Sberbank",
			Description: fmt.Sprintf("Credit card of %s", username),
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		}

		err := notesUseCase.UpdateCard(context.Background(), card, true)

		assert.NoError(t, err)
	})

	t.Run("should return undated note by identifier", func(t *testing.T) {
		card, err := notesUseCase.GetCard(context.Background(), id)

		expected := &entity.Card{
			Title:       "Sberbank",
			Description: fmt.Sprintf("Credit card of %s", username),
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			Version:     2,
		}

		assert.Equal(t, expected, card)
		assert.NoError(t, err)
	})

	t.Run("should successfully delete the note", func(t *testing.T) {
		err := notesUseCase.DeleteCard(context.Background(), id)

		assert.NoError(t, err)
	})

	t.Run("should return an empty list", func(t *testing.T) {
		list, err := notesUseCase.GetCardsList(context.Background())

		assert.Empty(t, list)
		assert.NoError(t, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		creds, err := notesUseCase.GetCard(context.Background(), id)

		assert.Nil(t, creds)
		assert.Equal(t, notes.ErrNotFound, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		card := &entity.Card{
			ID:          id,
			Title:       "Sberbank",
			Description: fmt.Sprintf("Credit card of %s", username),
			Number:      "1234 5678 9012 4453",
			Holder:      "John Daw",
			CVC:         "123",
			Expiry:      time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
		}

		err := notesUseCase.UpdateCard(context.Background(), card, true)

		assert.Equal(t, notes.ErrNotFound, err)
	})

	t.Run("should fail with nothing found error", func(t *testing.T) {
		err := notesUseCase.DeleteCard(context.Background(), id)

		assert.Equal(t, notes.ErrNotFound, err)
	})
}
