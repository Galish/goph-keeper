package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) UpdateCredentials(ctx context.Context, creds *entity.Credentials) error {
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

		CreatedBy:    creds.CreatedBy,
		LastEditedAt: time.Now(),
	}

	err := uc.repo.UpdateSecureRecord(ctx, record)
	if errors.Is(err, repository.ErrNothingFound) {
		return ErrNothingFound
	}

	return err
}
