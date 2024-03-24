package psql

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) SetSecureRecord(ctx context.Context, record *repository.SecureRecord) error {
	return nil
}
