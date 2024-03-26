package psql

import (
	"database/sql"
	"errors"
	"os"

	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	errCodeConflict       = "23505"
	errCodeCheckViolation = "23514"
)

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

	if err := store.Bootstrap(cfg.DBInitPath); err != nil {
		logger.WithError(err).Debug("database initialization error")
	}

	return &store, nil
}

func (s *psqlStore) Bootstrap(filePath string) error {
	logger.Info("database initialization")

	query, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(string(query))
	if err != nil {
		return err
	}

	return nil
}

func (s *psqlStore) Close() error {
	return s.db.Close()
}
