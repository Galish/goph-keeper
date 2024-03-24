package psql

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) GetSecureRecord(ctx context.Context, user, id string) (*repository.SecureRecord, error) {
	return nil, nil
}
