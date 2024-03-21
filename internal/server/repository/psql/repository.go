package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// const (
// 	errCodeConflict       = "23505"
// 	errCodeCheckViolation = "23514"
// )

type psqlStore struct {
	db *sql.DB
}

func New(cfg *config.Config) (*psqlStore, error) {
	if cfg.DBAddr == "" {
		return nil, errors.New("database address missing")
	}

	logger.Info("database connection")

	db, err := sql.Open("pgx", cfg.DBAddr)
	if err != nil {
		return nil, err
	}

	store := psqlStore{db}

	logger.Info("database initialization")

	// if err := store.init(); err != nil {
	// 	return nil, err
	// }

	return &store, nil
}

// func (s *psqlStore) init() error {
// 	query, err := os.ReadFile("internal/app/adapters/repository/psql/init.sql")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = s.db.Exec(string(query))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (s *psqlStore) CreateUser(ctx context.Context, user *entity.User) error {
	return nil
}

func (s *psqlStore) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	return nil, nil
}

func (s *psqlStore) CreateSecureRecord(ctx context.Context, record *repository.SecureRecord) error {
	return nil
}

func (s *psqlStore) GetSecureRecord(ctx context.Context, user, id string) (*repository.SecureRecord, error) {
	return nil, nil
}

func (s *psqlStore) GetSecureRecords(ctx context.Context, user string, recordType repository.SecureRecordType) ([]*repository.SecureRecord, error) {
	return nil, nil
}

func (s *psqlStore) Close() error {
	return s.db.Close()
}
