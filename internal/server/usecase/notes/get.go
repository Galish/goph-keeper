package notes

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *NotesUseCase) Get(ctx context.Context, id string) (*entity.Note, error) {
	record, err := uc.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if record.Type != repository.TypeNote {
		return nil, ErrInvalidType
	}

	note := entity.Note{
		ID:           record.ID,
		Title:        record.Title,
		Description:  record.Description,
		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
	}

	if record.TextNote != "" {
		note.Value = record.TextNote
	} else if record.RawNote != nil {
		note.RawValue = record.RawNote
	}

	return &note, nil
}
