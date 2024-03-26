package keeper

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) GetTextNotes(
	ctx context.Context,
	user string,
) ([]*entity.Note, error) {
	records, err := uc.repo.GetSecureRecords(ctx, user, repository.TypeNote)
	if err != nil {
		return nil, err
	}

	var notes = make([]*entity.Note, len(records))

	for i, r := range records {
		note := &entity.Note{
			ID:           r.ID,
			Title:        r.Title,
			Description:  r.Description,
			Value:        r.TextNote,
			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		notes[i] = note
	}

	return notes, nil
}

func (uc *KeeperUseCase) GetRawNotes(
	ctx context.Context,
	user string,
) ([]*entity.Note, error) {
	records, err := uc.repo.GetSecureRecords(ctx, user, repository.TypeRawNote)
	if err != nil {
		return nil, err
	}

	var notes = make([]*entity.Note, len(records))

	for i, r := range records {
		note := &entity.Note{
			ID:           r.ID,
			Title:        r.Title,
			Description:  r.Description,
			RawValue:     r.RawNote,
			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		notes[i] = note
	}

	return notes, nil
}

func (uc *KeeperUseCase) GetCards(
	ctx context.Context,
	user string,
) ([]*entity.Card, error) {
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

func (uc *KeeperUseCase) GetAllCredentials(
	ctx context.Context,
	user string,
) ([]*entity.Credentials, error) {
	records, err := uc.repo.GetSecureRecords(ctx, user, repository.TypeCredentials)
	if err != nil {
		return nil, err
	}

	var creds = make([]*entity.Credentials, len(records))

	for i, r := range records {
		cred := &entity.Credentials{
			ID:          r.ID,
			Title:       r.Title,
			Description: r.Description,

			Username: r.Username,
			Password: r.Password,

			CreatedAt:    r.CreatedAt,
			LastEditedAt: r.LastEditedAt,
		}

		creds[i] = cred
	}

	return creds, nil
}
