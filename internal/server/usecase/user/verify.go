package user

import "github.com/Galish/goph-keeper/internal/entity"

func (uc *UserUseCase) Verify(accessToken string) (*entity.User, error) {
	claims, err := uc.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID: claims.UserID,
	}

	return user, nil
}
