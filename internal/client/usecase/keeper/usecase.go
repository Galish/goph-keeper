package keeper

import (
	pb "github.com/Galish/goph-keeper/api/proto"
)

type KeeperUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *KeeperUseCase {
	return &KeeperUseCase{
		client: client,
	}
}
