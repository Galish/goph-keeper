package notes

import (
	"context"
	"errors"
	"time"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *UseCase) AddCredentials(ctx context.Context, creds *entity.Credentials) error {
	if creds == nil || !creds.IsValid() {
		return ErrInvalidEntity
	}

	note := &repository.SecureNote{
		ID:          creds.ID,
		Type:        repository.TypeCredentials,
		Title:       creds.Title,
		Description: creds.Description,

		Username: creds.Username,
		Password: creds.Password,

		CreatedBy: creds.CreatedBy,
		CreatedAt: time.Now(),
	}

	return uc.repo.AddSecureNote(ctx, note)
}

func (uc *UseCase) GetCredentials(ctx context.Context, user, id string) (*entity.Credentials, error) {
	if id == "" || user == "" {
		return nil, ErrMissingArgument
	}

	note, err := uc.repo.GetSecureNote(ctx, user, id, repository.TypeCredentials)
	if errors.Is(err, repository.ErrNotFound) {
		return nil, ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	creds := entity.Credentials{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,

		Username: note.Username,
		Password: note.Password,

		CreatedAt:    note.CreatedAt,
		LastEditedAt: note.LastEditedAt,
		Version:      note.Version,
	}

	return &creds, nil
}

func (uc *UseCase) GetAllCredentials(ctx context.Context, user string) ([]*entity.Credentials, error) {
	if user == "" {
		return nil, ErrMissingArgument
	}

	notes, err := uc.repo.GetSecureNotes(ctx, user, repository.TypeCredentials)
	if err != nil {
		return nil, err
	}

	var creds = make([]*entity.Credentials, len(notes))

	for i, r := range notes {
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

func (uc *UseCase) UpdateCredentials(ctx context.Context, creds *entity.Credentials, overwrite bool) error {
	if creds == nil || creds.ID == "" || !creds.IsValid() {
		return ErrInvalidEntity
	}

	if !overwrite && creds.Version == 0 {
		return ErrVersionRequired
	}

	note := &repository.SecureNote{
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
		note.Version = creds.Version
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

func (uc *UseCase) DeleteCredentials(ctx context.Context, user, id string) error {
	if id == "" || user == "" {
		return ErrMissingArgument
	}

	err := uc.repo.DeleteSecureNote(ctx, user, id, repository.TypeCredentials)
	if errors.Is(err, repository.ErrNotFound) {
		return ErrNotFound
	}

	return err
}
