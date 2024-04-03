package keeper

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"github.com/Galish/goph-keeper/internal/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *KeeperUseCase) AddCredentials(creds *entity.Credentials) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.AddCredentialsRequest{
		Credentials: &pb.Credentials{
			Title:       creds.Title,
			Description: creds.Description,
			Username:    creds.Username,
			Password:    creds.Password,
		},
	}

	_, err := uc.client.AddCredentials(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) UpdateCredentials(creds *entity.Credentials, overwrite bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.UpdateCredentialsRequest{
		Id: creds.ID,
		Credentials: &pb.Credentials{
			Title:       creds.Title,
			Description: creds.Description,
			Username:    creds.Username,
			Password:    creds.Password,
		},
		Version:   creds.Version,
		Overwrite: overwrite,
	}

	_, err := uc.client.UpdateCredentials(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) GetCredentials(id string) (*entity.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetCredentials(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	creds := &entity.Credentials{
		Title:       resp.Credentials.GetTitle(),
		Description: resp.Credentials.GetDescription(),
		Username:    resp.Credentials.GetUsername(),
		Password:    resp.Credentials.GetPassword(),
		Version:     resp.GetVersion(),
	}

	return creds, nil
}

func (uc *KeeperUseCase) GetCredentialsList() ([]*entity.Credentials, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	resp, err := uc.client.GetCredentialsList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, handleError(err)
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

func (uc *KeeperUseCase) DeleteCredentials(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteCredentials(ctx, req)

	return handleError(err)
}
