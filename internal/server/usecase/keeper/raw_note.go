package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) AddRawNote(ctx context.Context, note *entity.RawNote) error {
	if note == nil || !note.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.SecureRecord{
		ID:          note.ID,
		Type:        repository.TypeRawNote,
		Title:       note.Title,
		Description: note.Description,

		RawNote: note.Value,

		CreatedBy: note.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureRecord(ctx, record)
}

func (uc *KeeperUseCase) GetRawNote(ctx context.Context, user, id string) (*entity.RawNote, error) {
	record, err := uc.repo.GetSecureRecord(ctx, user, id, repository.TypeRawNote)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	note := entity.RawNote{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,

		Value: record.RawNote,

		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
		Version:      record.Version,
	}

	return &note, nil
}

func (uc *KeeperUseCase) GetRawNotes(ctx context.Context, user string) ([]*entity.RawNote, error) {
	records, err := uc.repo.GetSecureRecords(ctx, user, repository.TypeRawNote)
	if err != nil {
		return nil, err
	}

	var notes = make([]*entity.RawNote, len(records))

	for i, r := range records {
		note := &entity.RawNote{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,

			Value: r.RawNote,

			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		notes[i] = note
	}

	return notes, nil
}

func (uc *KeeperUseCase) UpdateRawNote(ctx context.Context, note *entity.RawNote, overwrite bool) error {
	if !overwrite && note.Version == 0 {
		return ErrVersionRequired
	}

	if note == nil || note.ID == "" || !note.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.SecureRecord{
		ID:          note.ID,
		Type:        repository.TypeRawNote,
		Title:       note.Title,
		Description: note.Description,

		RawNote: note.Value,

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

func (uc *KeeperUseCase) DeleteRawNote(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeRawNote)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
