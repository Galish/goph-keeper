// Package encryption provides functions for data encryption.
package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// Encryptor the authentication manager.
type Encryptor interface {
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

// AESEncryptor represents the AES encryptor.
type AESEncryptor struct {
	passphrase []byte
}

// NewAESEncryptor returns a new AES encryptor instance.
func NewAESEncryptor(passphrase []byte) *AESEncryptor {
	return &AESEncryptor{
		passphrase: passphrase,
	}
}

// Encrypt encrypts a string using the AES algorithm.
func (e *AESEncryptor) Encrypt(input string) (string, error) {
	block, err := aes.NewCipher(e.passphrase)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encrypted := gcm.Seal(nonce, nonce, []byte(input), nil)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts a string using the AES algorithm.
func (e *AESEncryptor) Decrypt(input string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(e.passphrase)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, cipher := str[:nonceSize], str[nonceSize:]

	b, err := gcm.Open(nil, nonce, cipher, nil)
	if err != nil {
		return "", err
	}

	return string(b), err
}
