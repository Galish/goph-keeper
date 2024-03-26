package keeper

import (
	"context"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Galish/goph-keeper/internal/client/entity"
)

type KeeperUseCase struct {
	client pb.KeeperClient
}

func New(client pb.KeeperClient) *KeeperUseCase {
	return &KeeperUseCase{
		client: client,
	}
}

func (uc *KeeperUseCase) GetCredentials(id string) (*entity.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetCredentials(ctx, req)
	if err != nil {
		return nil, err
	}

	creds := &entity.Credentials{
		ID:          resp.Credentials.GetPassword(),
		Title:       resp.Credentials.GetTitle(),
		Description: resp.Credentials.GetDescription(),
		Username:    resp.Credentials.GetUsername(),
		Password:    resp.Credentials.GetPassword(),
	}

	return creds, nil
}

func (uc *KeeperUseCase) DeleteCredentials(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteCredentials(ctx, req)
	if err != nil {
		return err
	}

	return nil
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
			ID:          c.GetId(),
			Title:       c.GetTitle(),
			Description: c.GetDescription(),
			Username:    c.GetUsername(),
			Password:    c.GetPassword(),
		}
	}

	return creds, nil
}
