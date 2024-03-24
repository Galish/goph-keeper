package psql

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) GetSecureRecords(ctx context.Context, user string, recordType repository.SecureRecordType) ([]*repository.SecureRecord, error) {
	return nil, nil
}
