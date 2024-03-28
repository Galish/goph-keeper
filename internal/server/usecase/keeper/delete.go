package keeper

import (
	"context"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *KeeperUseCase) DeleteTextNote(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeNote)
	if errors.Is(err, repository.ErrNothingFound) {
		return ErrNothingFound
	}
	if err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) DeleteRawNote(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeRawNote)
	if errors.Is(err, repository.ErrNothingFound) {
		return ErrNothingFound
	}
	if err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) DeleteCard(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeCard)
	if errors.Is(err, repository.ErrNothingFound) {
		return ErrNothingFound
	}
	if err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) DeleteCredentials(ctx context.Context, user, id string) error {
	err := uc.repo.DeleteSecureRecord(ctx, user, id, repository.TypeCredentials)
	if errors.Is(err, repository.ErrNothingFound) {
		return ErrNothingFound
	}
	if err != nil {
		return err
	}

	return nil
}
