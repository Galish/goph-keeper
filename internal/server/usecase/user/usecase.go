package user

import (
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/pkg/auth"
)

type UserUseCase struct {
	repo       repository.UserRepository
	jwtManager *auth.JWTManager
}

func New(repo repository.UserRepository, jwtManager *auth.JWTManager) *UserUseCase {
	return &UserUseCase{
		repo:       repo,
		jwtManager: jwtManager,
	}
}
