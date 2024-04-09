package user

import (
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var defaultTimeout = 1 * time.Minute

type UseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *UseCase {
	return &UseCase{
		client: client,
	}
}
