package notes

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Galish/goph-keeper/internal/entity"
)

// AddCredentials creates a new login password pair.
func (uc *UseCase) AddCredentials(ctx context.Context, creds *entity.Credentials) error {
	if creds == nil || !creds.IsValid() {
		return ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

// UpdateCredentials updates a login password pair.
func (uc *UseCase) UpdateCredentials(ctx context.Context, creds *entity.Credentials, overwrite bool) error {
	if creds == nil || creds.ID == "" || !creds.IsValid() {
		return ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

// GetCredentials returns a login password pair for the given identifier.
func (uc *UseCase) GetCredentials(ctx context.Context, id string) (*entity.Credentials, error) {
	if id == "" {
		return nil, ErrMissingArgument
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetCredentials(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	creds := &entity.Credentials{
		Title:       resp.GetCredentials().GetTitle(),
		Description: resp.GetCredentials().GetDescription(),
		Username:    resp.GetCredentials().GetUsername(),
		Password:    resp.GetCredentials().GetPassword(),
		Version:     resp.GetVersion(),
	}

	return creds, nil
}

// GetCredentialsList returns login password pairs list.
func (uc *UseCase) GetCredentialsList(ctx context.Context) ([]*entity.Credentials, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

// DeleteCredentials deletes the login password pair for the given identifier.
func (uc *UseCase) DeleteCredentials(ctx context.Context, id string) error {
	if id == "" {
		return ErrMissingArgument
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteCredentials(ctx, req)

	return handleError(err)
}
