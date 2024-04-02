package auth

import (
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/entity"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID string
}

type JWTManager struct {
	secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
	return &JWTManager{
		secretKey,
	}
}

func (m *JWTManager) Generate(user *entity.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&JWTClaims{
			UserID: user.ID,
		},
	)

	return token.SignedString([]byte(m.secretKey))
}

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
