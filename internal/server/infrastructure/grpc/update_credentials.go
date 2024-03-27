package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) UpdateCredentials(
	ctx context.Context,
	in *pb.UpdateCredentialsRequest,
) (*emptypb.Empty, error) {
	creds := &entity.Credentials{
		ID:          in.Credentials.GetId(),
		Title:       in.Credentials.GetTitle(),
		Description: in.Credentials.GetDescription(),
		Username:    in.Credentials.GetUsername(),
		Password:    in.Credentials.GetPassword(),
		CreatedBy:   ctx.Value(interceptors.UserContextKey).(string),
	}

	err := s.keeper.UpdateCredentials(ctx, creds)
	if errors.Is(err, keeper.ErrNothingFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return nil, nil
}
