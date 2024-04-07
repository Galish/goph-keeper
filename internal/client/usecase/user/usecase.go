package user

import (
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var defaultTimeout = 1 * time.Minute

type UserUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *UserUseCase {
	return &UserUseCase{
		client: client,
	}
}
