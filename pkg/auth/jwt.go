// Package auth provides functions for user authentication and authorization.
package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// JWTClaims represents data encoded into a token.
type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string
}

// JWTManager represents the authentication manager.
type JWTManager struct {
	secretKey string
}

// NewJWTManager returns a new JWT manager instance.
func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{
		secretKey,
	}
}

// Generate creates and returns a JWT token string.
func (m *JWTManager) Generate(claims *JWTClaims) (string, error) {
	if claims == nil {
		return "", errors.New("missing required argument")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString([]byte(m.secretKey))
}

// Verify decodes JWT token string.
func (m *JWTManager) Verify(accessToken string) (*JWTClaims, error) {
	var claims JWTClaims

	token, err := jwt.ParseWithClaims(
		accessToken,
		&claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}

			return []byte(m.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token is not valid")
	}

	return &claims, nil
}
