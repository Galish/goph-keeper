package keeper

import (
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
)

var defaultTimeout = 1 * time.Minute

type KeeperUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *KeeperUseCase {
	return &KeeperUseCase{
		client: client,
	}
}
