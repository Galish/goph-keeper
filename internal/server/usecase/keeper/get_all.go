package keeper

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) GetAllNotes(ctx context.Context) ([]*entity.Note, error) {
	records, err := uc.repo.GetByType(ctx, repository.TypeNote)
	if err != nil {
		return nil, err
	}

	var notes = make([]*entity.Note, len(records))

	for i, r := range records {
		note := &entity.Note{
			ID:           r.ID,
			Title:        r.Title,
			Description:  r.Description,
			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		if r.TextNote != "" {
			note.Value = r.TextNote
		} else if r.RawNote != nil {
			note.RawValue = r.RawNote
		}

		notes[i] = note
	}

	return notes, nil
}
