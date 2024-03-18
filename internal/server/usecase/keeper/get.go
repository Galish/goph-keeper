package keeper

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) GetNote(
	ctx context.Context,
	user, id string,
) (*entity.Note, error) {
	record, err := uc.repo.Get(ctx, user, id)
	if err != nil {
		return nil, err
	}

	if record.Type != repository.TypeNote &&
		record.Type != repository.TypeRawNote {
		return nil, ErrInvalidType
	}

	note := entity.Note{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,

		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
	}

	if record.Type == repository.TypeNote {
		note.Value = record.TextNote
	} else if record.Type == repository.TypeRawNote {
		note.RawValue = record.RawNote
	}

	return &note, nil
}

func (uc *KeeperUseCase) GetCard(
	ctx context.Context,
	user, id string,
) (*entity.Card, error) {
	record, err := uc.repo.Get(ctx, user, id)
	if err != nil {
		return nil, err
	}

	if record.Type != repository.TypeCard {
		return nil, ErrInvalidType
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

func (uc *KeeperUseCase) GetCredentials(
	ctx context.Context,
	user, id string,
) (*entity.Credentials, error) {
	record, err := uc.repo.Get(ctx, user, id)
	if err != nil {
		return nil, err
	}

	if record.Type != repository.TypeCreds {
		return nil, ErrInvalidType
	}

	creds := entity.Credentials{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,

		Username: record.Username,
		Password: record.Password,

		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
	}

	return &creds, nil
}
