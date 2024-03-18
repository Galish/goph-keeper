package notes

import (
	"context"
	"time"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *NotesUseCase) Add(ctx context.Context, note *entity.Note) error {
	if note == nil || !note.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.Record{
		ID:          note.ID,
		Type:        repository.TypeNote,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   time.Now(),
	}

	if note.Value != "" {
		record.TextNote = note.Value
	} else if string(note.RawValue) != "" {
		record.RawNote = note.RawValue
	}

	return uc.repo.Set(ctx, record)
}
