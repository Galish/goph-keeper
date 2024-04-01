package user

import (
	pb "github.com/Galish/goph-keeper/api/proto"
)

type UserUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *UserUseCase {
	return &UserUseCase{
		client: client,
	}
}
