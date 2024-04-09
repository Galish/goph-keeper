package notes

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *UseCase) AddRawNote(ctx context.Context, rawNote *entity.RawNote) error {
	if rawNote == nil || !rawNote.IsValid() {
		return ErrInvalidEntity
	}

	note := &repository.SecureNote{
		ID:          rawNote.ID,
		Type:        repository.TypeRawNote,
		Title:       rawNote.Title,
		Description: rawNote.Description,

		RawNote: rawNote.Value,

		CreatedBy: rawNote.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureNote(ctx, note)
}

func (uc *UseCase) GetRawNote(ctx context.Context, user, id string) (*entity.RawNote, error) {
	if id == "" || user == "" {
		return nil, ErrMissingArgument
	}

	note, err := uc.repo.GetSecureNote(ctx, user, id, repository.TypeRawNote)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	rawNote := entity.RawNote{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,

		Value: note.RawNote,

		CreatedAt:    note.CreatedAt,
		LastEditedAt: note.LastEditedAt,
		Version:      note.Version,
	}

	return &rawNote, nil
}

func (uc *UseCase) GetRawNotes(ctx context.Context, user string) ([]*entity.RawNote, error) {
	if user == "" {
		return nil, ErrMissingArgument
	}

	notes, err := uc.repo.GetSecureNotes(ctx, user, repository.TypeRawNote)
	if err != nil {
		return nil, err
	}

	var rawNotes = make([]*entity.RawNote, len(notes))

	for i, r := range notes {
		note := &entity.RawNote{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,

			Value: r.RawNote,

			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		rawNotes[i] = note
	}

	return rawNotes, nil
}

func (uc *UseCase) UpdateRawNote(ctx context.Context, rawNote *entity.RawNote, overwrite bool) error {
	if rawNote == nil || rawNote.ID == "" || !rawNote.IsValid() {
		return ErrInvalidEntity
	}

	if !overwrite && rawNote.Version == 0 {
		return ErrVersionRequired
	}

	note := &repository.SecureNote{
		ID:          rawNote.ID,
		Type:        repository.TypeRawNote,
		Title:       rawNote.Title,
		Description: rawNote.Description,

		RawNote: rawNote.Value,

		CreatedBy:    rawNote.CreatedBy,
		LastEditedAt: time.Now(),
	}

	if !overwrite {
		note.Version = rawNote.Version
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

func (uc *UseCase) DeleteRawNote(ctx context.Context, user, id string) error {
	if id == "" || user == "" {
		return ErrMissingArgument
	}

	err := uc.repo.DeleteSecureNote(ctx, user, id, repository.TypeRawNote)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
