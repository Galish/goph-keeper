package keeper

import (
	"context"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/client/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *KeeperUseCase) AddCredentials(creds *entity.Credentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.AddCredentialsRequest{
		Credentials: &pb.Credentials{
			Title:       creds.Title,
			Description: creds.Description,
			Username:    creds.Username,
			Password:    creds.Password,
		},
	}

	if _, err := uc.client.AddCredentials(ctx, req); err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) UpdateCredentials(creds *entity.Credentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.UpdateCredentialsRequest{
		Id: creds.ID,
		Credentials: &pb.Credentials{
			Title:       creds.Title,
			Description: creds.Description,
			Username:    creds.Username,
			Password:    creds.Password,
		},
	}

	if _, err := uc.client.UpdateCredentials(ctx, req); err != nil {
		return err
	}

	return nil
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

	resp, err := uc.client.GetCredentialsList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var creds = make([]*entity.Credentials, len(resp.GetList()))

	for i, c := range resp.GetList() {
		creds[i] = &entity.Credentials{
			ID:          c.GetId(),
			Title:       c.GetTitle(),
			Description: c.GetDescription(),
		}
	}

	return creds, nil
}
