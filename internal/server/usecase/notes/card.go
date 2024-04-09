package notes

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *UseCase) AddCard(ctx context.Context, card *entity.Card) error {
	if card == nil || !card.IsValid() {
		return ErrInvalidEntity
	}

	note := &repository.SecureNote{
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

	return uc.repo.AddSecureNote(ctx, note)
}

func (uc *UseCase) GetCard(ctx context.Context, user, id string) (*entity.Card, error) {
	if id == "" || user == "" {
		return nil, ErrMissingArgument
	}

	note, err := uc.repo.GetSecureNote(ctx, user, id, repository.TypeCard)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	card := entity.Card{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,

		Number: note.CardNumber,
		Holder: note.CardHolder,
		CVC:    note.CardCVC,
		Expiry: note.CardExpiry,

		CreatedAt:    note.CreatedAt,
		LastEditedAt: note.LastEditedAt,
		Version:      note.Version,
	}

	return &card, nil
}

func (uc *UseCase) GetCards(ctx context.Context, user string) ([]*entity.Card, error) {
	if user == "" {
		return nil, ErrMissingArgument
	}

	notes, err := uc.repo.GetSecureNotes(ctx, user, repository.TypeCard)
	if err != nil {
		return nil, err
	}

	var cards = make([]*entity.Card, len(notes))

	for i, r := range notes {
		card := &entity.Card{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,

			Number: r.CardNumber,
			Holder: r.CardHolder,
			CVC:    r.CardCVC,
			Expiry: r.CardExpiry,

			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		cards[i] = card
	}

	return cards, nil
}

func (uc *UseCase) UpdateCard(ctx context.Context, card *entity.Card, overwrite bool) error {
	if card == nil || card.ID == "" || !card.IsValid() {
		return ErrInvalidEntity
	}

	if !overwrite && card.Version == 0 {
		return ErrVersionRequired
	}

	note := &repository.SecureNote{
		ID:          card.ID,
		Type:        repository.TypeCard,
		Title:       card.Title,
		Description: card.Description,

		CardNumber: card.Number,
		CardHolder: card.Holder,
		CardCVC:    card.CVC,
		CardExpiry: card.Expiry,

		CreatedBy:    card.CreatedBy,
		LastEditedAt: time.Now(),
	}

	if !overwrite {
		note.Version = card.Version
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

func (uc *UseCase) DeleteCard(ctx context.Context, user, id string) error {
	if id == "" || user == "" {
		return ErrMissingArgument
	}

	err := uc.repo.DeleteSecureNote(ctx, user, id, repository.TypeCard)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
