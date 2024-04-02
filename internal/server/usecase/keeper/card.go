package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

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

func (uc *KeeperUseCase) GetCard(ctx context.Context, user, id string) (*entity.Card, error) {
	record, err := uc.repo.GetSecureRecord(ctx, user, id, repository.TypeCard)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	card := entity.Card{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,

		Number: record.CardNumber,
		Holder: record.CardHolder,
		CVC:    record.CardCVC,
		Expiry: record.CardExpiry,

		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
	}

	return &card, nil
}

func (uc *KeeperUseCase) GetCards(ctx context.Context, user string) ([]*entity.Card, error) {
	records, err := uc.repo.GetSecureRecords(ctx, user, repository.TypeCard)
	if err != nil {
		return nil, err
	}

	var cards = make([]*entity.Card, len(records))

	for i, r := range records {
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

func (uc *KeeperUseCase) UpdateCard(ctx context.Context, card *entity.Card) error {
	if card == nil || card.ID == "" || !card.IsValid() {
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

		CreatedBy:    card.CreatedBy,
		LastEditedAt: time.Now(),
	}

	err := uc.repo.UpdateSecureRecord(ctx, record)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}

func (uc *KeeperUseCase) DeleteCard(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeCard)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
