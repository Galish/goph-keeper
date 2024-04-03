package keeper

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

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

func (uc *KeeperUseCase) GetCredentials(ctx context.Context, user, id string) (*entity.Credentials, error) {
	record, err := uc.repo.GetSecureRecord(ctx, user, id, repository.TypeCredentials)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	creds := entity.Credentials{
		ID:          record.ID,
		Title:       record.Title,
		Description: record.Description,

		Username: record.Username,
		Password: record.Password,

		CreatedAt:    record.CreatedAt,
		LastEditedAt: record.LastEditedAt,
		Version:      record.Version,
	}

	return &creds, nil
}

func (uc *KeeperUseCase) GetAllCredentials(ctx context.Context, user string) ([]*entity.Credentials, error) {
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

func (uc *KeeperUseCase) UpdateCredentials(ctx context.Context, creds *entity.Credentials, overwrite bool) error {
	if !overwrite && creds.Version == 0 {
		return ErrVersionRequired
	}

	if creds == nil || creds.ID == "" || !creds.IsValid() {
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

	if !overwrite {
		record.Version = creds.Version
	}

	err := uc.repo.UpdateSecureRecord(ctx, record)
	if errors.Is(err, repository.ErrVersionConflict) {
		return ErrVersionConflict
	}

	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}

func (uc *KeeperUseCase) DeleteCredentials(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeCredentials)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
