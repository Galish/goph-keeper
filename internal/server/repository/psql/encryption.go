package psql

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

type protectedData struct {
	Username   string
	Password   string
	TextNote   string
	RawNote    []byte
	CardNumber string
	CardHolder string
	CardCVC    string
	CardExpiry time.Time
}

func (s *Store) encrypt(r *repository.SecureNote) (string, error) {
	b, err := json.Marshal(protectedData{
		Username:   r.Username,
		Password:   r.Password,
		TextNote:   r.TextNote,
		RawNote:    r.RawNote,
		CardNumber: r.CardNumber,
		CardHolder: r.CardHolder,
		CardCVC:    r.CardCVC,
		CardExpiry: r.CardExpiry,
	})
	if err != nil {
		return "", fmt.Errorf("failed to encode data for encryption: %w", err)
	}

	encrypted, err := s.enc.Encrypt(string(b))
	if err != nil {
		return "", fmt.Errorf("failed data encryption: %w", err)
	}

	return encrypted, err
}

func (s *Store) decrypt(raw string, r *repository.SecureNote) error {
	decrypted, err := s.enc.Decrypt(raw)
	if err != nil {
		return fmt.Errorf("failed data decryption: %w", err)
	}

	var protected = new(protectedData)
	if err := json.Unmarshal([]byte(decrypted), protected); err != nil {
		return fmt.Errorf("failed to decode decrypted data: %w", err)
	}

	r.Username = protected.Username
	r.Password = protected.Password
	r.TextNote = protected.TextNote
	r.RawNote = protected.RawNote
	r.CardNumber = protected.CardNumber
	r.CardHolder = protected.CardHolder
	r.CardCVC = protected.CardCVC
	r.CardExpiry = protected.CardExpiry

	return nil
}
