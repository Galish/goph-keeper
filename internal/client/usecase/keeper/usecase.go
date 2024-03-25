package keeper

import (
	"context"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

type KeeperUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *KeeperUseCase {
	return &KeeperUseCase{
		client: client,
	}
}

func (uc *KeeperUseCase) GetAllCredentials() ([]*entity.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	resp, err := uc.client.GetAllCredentials(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var creds = make([]*entity.Credentials, len(resp.Credentials))

	for i, c := range resp.Credentials {
		creds[i] = &entity.Credentials{
			Title:       c.Title,
			Description: c.Description,
			Username:    c.Username,
			Password:    c.Password,
		}
	}

	return creds, nil
}
