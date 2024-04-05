package psql

import (
	"encoding/json"
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

func (s *psqlStore) encrypt(r *repository.SecureRecord) (string, error) {
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
		return "", err
	}

	return s.enc.Encrypt(string(b))
}

func (s *psqlStore) decrypt(raw string, r *repository.SecureRecord) error {
	decrypted, err := s.enc.Decrypt(raw)
	if err != nil {
		return nil
	}

	var protected = new(protectedData)
	if err := json.Unmarshal([]byte(decrypted), protected); err != nil {
		return err
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
