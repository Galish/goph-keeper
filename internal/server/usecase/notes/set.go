package notes

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *NotesUseCase) Set(ctx context.Context, note *entity.Note) error {
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
	} else {
		return errors.New("value is missing")
	}

	return uc.repo.Set(ctx, record)
}
