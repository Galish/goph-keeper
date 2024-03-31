package user

import (
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var (
	ErrAlreadyExists      = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("incorrect login/password pair")
	ErrNoPassword         = errors.New("password not specified")
	ErrNoUsername         = errors.New("username not specified")
	ErrNotFound           = errors.New("user not found")
)

type UserUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *UserUseCase {
	return &UserUseCase{
		client: client,
	}
}
