package user

import (
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/pkg/auth"
)

type UseCase struct {
	repo       repository.UserRepository
	jwtManager *auth.JWTManager
}

func New(repo repository.UserRepository, jwtManager *auth.JWTManager) *UseCase {
	return &UseCase{
		repo:       repo,
		jwtManager: jwtManager,
	}
}
