package user

import (
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var (
	ErrInvalidCredentials = errors.New("incorrect login/password pair")
	ErrNoPassword         = errors.New("password not specified")
	ErrNoUsername         = errors.New("username not specified")
)

type UserUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *UserUseCase {
	return &UserUseCase{
		client: client,
	}
}
