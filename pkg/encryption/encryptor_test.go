package encryption_test

import (
	"errors"
	"testing"

	"github.com/Galish/goph-keeper/pkg/encryption"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAESEncryptor(t *testing.T) {
	tests := []struct {
		name       string
		passphrase string
		input      string
		err        error
	}{
		{
			"invalid passphrase size",
			"pqssjyEpfbwy",
			"",
			errors.New("crypto/aes: invalid key size 12"),
		},
		{
			"empty input",
			"pqssjyEpfbwxyAqTPJdP28ueaVmrjEjV",
			"",
			nil,
		},
		{
			"input string",
			"pqssjyEpfbwxyAqTPJdP28ueaVmrjEjV",
			"Hello world!",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			enc := encryption.NewAESEncryptor([]byte(tt.passphrase))

			encrypted, err := enc.Encrypt(tt.input)
			if err != nil {
				require.ErrorContains(t, tt.err, err.Error())
			}

			decrypted, err := enc.Decrypt(encrypted)
			if err != nil {
				require.ErrorContains(t, tt.err, err.Error())
			}

			assert.Equal(t, tt.input, decrypted)
		})
	}
}
