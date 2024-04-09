package psql

import (
	"database/sql"
	"errors"
	"os"

	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/pkg/encryption"
	"github.com/Galish/goph-keeper/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	errCodeConflict       = "23505"
	errCodeCheckViolation = "23514"
)

type Store struct {
	db  *sql.DB
	enc encryption.Encryptor
}

func New(cfg *config.Config) (*Store, error) {
	if cfg.DBAddr == "" {
		return nil, errors.New("database address missing")
	}

	logger.Info("database connection")

	db, err := sql.Open("pgx", cfg.DBAddr)
	if err != nil {
		return nil, err
	}

	store := &Store{
		db:  db,
		enc: encryption.NewAESEncryptor([]byte(cfg.EncryptPassphrase)),
	}

	if err := store.Bootstrap(cfg.DBInitPath); err != nil {
		logger.WithError(err).Fatal("database initialization error")
	}

	return store, nil
}

func (s *Store) Bootstrap(filePath string) error {
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

func (s *Store) Close() error {
	logger.Info("shutting down the DB server")

	return s.db.Close()
}
