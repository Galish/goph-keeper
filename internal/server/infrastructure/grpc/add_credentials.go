package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddCredentials(ctx context.Context, in *pb.AddCredentialsRequest) (*pb.AddCredentialsResponse, error) {
	var response pb.AddCredentialsResponse

	creds := entity.NewCredentials()
	creds.Title = in.Credentials.GetTitle()
	creds.Description = in.Credentials.GetDescription()
	creds.Username = in.Credentials.GetUsername()
	creds.Password = in.Credentials.GetPassword()
	creds.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	err := s.keeper.AddCredentials(ctx, creds)
	if errors.Is(err, keeper.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		logger.WithError(err).Debug("unable to add the record")

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Id = creds.ID

	return &response, nil
}
