package notes

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *UseCase) AddTextNote(ctx context.Context, textNote *entity.TextNote) error {
	if textNote == nil || !textNote.IsValid() {
		return ErrInvalidEntity
	}

	note := &repository.SecureNote{
		ID:          textNote.ID,
		Type:        repository.TypeTextNote,
		Title:       textNote.Title,
		Description: textNote.Description,

		TextNote: textNote.Value,

		CreatedBy: textNote.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureNote(ctx, note)
}

func (uc *UseCase) GetTextNote(ctx context.Context, user, id string) (*entity.TextNote, error) {
	if id == "" || user == "" {
		return nil, ErrMissingArgument
	}

	note, err := uc.repo.GetSecureNote(ctx, user, id, repository.TypeTextNote)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	textNote := entity.TextNote{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,

		Value: note.TextNote,

		CreatedAt:    note.CreatedAt,
		LastEditedAt: note.LastEditedAt,
		Version:      note.Version,
	}

	return &textNote, nil
}

func (uc *UseCase) GetTextNotes(ctx context.Context, user string) ([]*entity.TextNote, error) {
	if user == "" {
		return nil, ErrMissingArgument
	}

	notes, err := uc.repo.GetSecureNotes(ctx, user, repository.TypeTextNote)
	if err != nil {
		return nil, err
	}

	var textNotes = make([]*entity.TextNote, len(notes))

	for i, r := range notes {
		note := &entity.TextNote{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,

			Value: r.TextNote,

			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		textNotes[i] = note
	}

	return textNotes, nil
}

func (uc *UseCase) UpdateTextNote(ctx context.Context, textNote *entity.TextNote, overwrite bool) error {
	if textNote == nil || textNote.ID == "" || !textNote.IsValid() {
		return ErrInvalidEntity
	}

	if !overwrite && textNote.Version == 0 {
		return ErrVersionRequired
	}

	note := &repository.SecureNote{
		ID:          textNote.ID,
		Type:        repository.TypeTextNote,
		Title:       textNote.Title,
		Description: textNote.Description,

		TextNote: textNote.Value,

		CreatedBy:    textNote.CreatedBy,
		LastEditedAt: time.Now(),
	}

	if !overwrite {
		note.Version = textNote.Version
	}

	err := uc.repo.UpdateSecureNote(ctx, note)
	if errors.Is(err, repository.ErrVersionConflict) {
		return ErrVersionConflict
	}

	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}

func (uc *UseCase) DeleteTextNote(ctx context.Context, user, id string) error {
	if id == "" || user == "" {
		return ErrMissingArgument
	}

	err := uc.repo.DeleteSecureNote(ctx, user, id, repository.TypeTextNote)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
