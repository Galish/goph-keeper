package auth

import (
	"testing"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

func BenchmarkGenerateToken(b *testing.B) {
	var (
		secretKey = "jwt_token_secret_key"
		userID    = "395fd5f4-964d-4135-9a55-fbf91c4a1614"
	)

	for i := 0; i < b.N; i++ {
		GenerateToken(
			secretKey,
			&entity.User{ID: userID},
		)
	}
}

func BenchmarkParseToken(b *testing.B) {
	var (
		secretKey   = "jwt_token_secret_key"
		tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIzOTVmZDVmNC05NjRkLTQxMzUtOWE1NS1mYmY5MWM0YTE2MTQifQ.tpw_R_Y-YdA6OV3tDz-KhWNH3-8FX4oVFur_BhYbmyQ"
	)

	for i := 0; i < b.N; i++ {
		ParseToken(secretKey, tokenString)
	}
}
