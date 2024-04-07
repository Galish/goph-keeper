package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) AddTextNote(ctx context.Context, note *entity.TextNote) error {
	if note == nil || !note.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.SecureRecord{
		ID:          note.ID,
		Type:        repository.TypeTextNote,
		Title:       note.Title,
		Description: note.Description,

		TextNote: note.Value,

		CreatedBy: note.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureRecord(ctx, record)
}

func (uc *KeeperUseCase) GetTextNote(ctx context.Context, user, id string) (*entity.TextNote, error) {
	if id == "" || user == "" {
		return nil, ErrMissingArgument
	}

	record, err := uc.repo.GetSecureRecord(ctx, user, id, repository.TypeTextNote)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	note := entity.TextNote{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,

		Value: record.TextNote,

		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
		Version:      record.Version,
	}

	return &note, nil
}

func (uc *KeeperUseCase) GetTextNotes(ctx context.Context, user string) ([]*entity.TextNote, error) {
	if user == "" {
		return nil, ErrMissingArgument
	}

	records, err := uc.repo.GetSecureRecords(ctx, user, repository.TypeTextNote)
	if err != nil {
		return nil, err
	}

	var notes = make([]*entity.TextNote, len(records))

	for i, r := range records {
		note := &entity.TextNote{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,

			Value: r.TextNote,

			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		notes[i] = note
	}

	return notes, nil
}

func (uc *KeeperUseCase) UpdateTextNote(ctx context.Context, note *entity.TextNote, overwrite bool) error {
	if note == nil || note.ID == "" || !note.IsValid() {
		return ErrInvalidEntity
	}

	if !overwrite && note.Version == 0 {
		return ErrVersionRequired
	}

	record := &repository.SecureRecord{
		ID:          note.ID,
		Type:        repository.TypeTextNote,
		Title:       note.Title,
		Description: note.Description,

		TextNote: note.Value,

		CreatedBy:    note.CreatedBy,
		LastEditedAt: time.Now(),
	}

	if !overwrite {
		record.Version = note.Version
	}

	err := uc.repo.UpdateSecureRecord(ctx, record)
	if errors.Is(err, repository.ErrVersionConflict) {
		return ErrVersionConflict
	}

	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}

func (uc *KeeperUseCase) DeleteTextNote(ctx context.Context, user, id string) error {
	if id == "" || user == "" {
		return ErrMissingArgument
	}

	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeTextNote)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
