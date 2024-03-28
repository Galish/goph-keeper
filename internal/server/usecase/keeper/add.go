package keeper

import (
	"context"
	"time"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) AddNote(ctx context.Context, note *entity.Note) error {
	if note == nil || !note.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.SecureRecord{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		CreatedBy:   note.CreatedBy,
		CreatedAt:   time.Now(),
	}

	if note.Value != "" {
		record.Type = repository.TypeNote
		record.TextNote = note.Value
	} else if string(note.RawValue) != "" {
		record.Type = repository.TypeRawNote
		record.RawNote = note.RawValue
	}

	return uc.repo.AddSecureRecord(ctx, record)
}

func (uc *KeeperUseCase) AddCard(ctx context.Context, card *entity.Card) error {
	if card == nil || !card.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.SecureRecord{
		ID:          card.ID,
		Type:        repository.TypeCard,
		Title:       card.Title,
		Description: card.Description,

		CardNumber: card.Number,
		CardHolder: card.Holder,
		CardCVC:    card.CVC,
		CardExpiry: card.Expiry,

		CreatedBy: card.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureRecord(ctx, record)
}

func (uc *KeeperUseCase) AddCredentials(ctx context.Context, creds *entity.Credentials) error {
	if creds == nil || !creds.IsValid() {
		return ErrInvalidEntity
	}

	record := &repository.SecureRecord{
		ID:          creds.ID,
		Type:        repository.TypeCredentials,
		Title:       creds.Title,
		Description: creds.Description,

		Username: creds.Username,
		Password: creds.Password,

		CreatedBy: creds.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureRecord(ctx, record)
}
