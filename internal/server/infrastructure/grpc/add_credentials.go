package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddCredentials(
	ctx context.Context,
	in *pb.AddCredentialsRequest,
) (*pb.AddCredentialsResponse, error) {
	creds := entity.NewCredentials()
	creds.Title = in.GetCredentials().GetTitle()
	creds.Description = in.GetCredentials().GetDescription()
	creds.Username = in.GetCredentials().GetUsername()
	creds.Password = in.GetCredentials().GetPassword()
	creds.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	err := s.notes.AddCredentials(ctx, creds)
	if err != nil {
		logger.WithError(err).Error("unable to add credentials")
	}

	if errors.Is(err, notes.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.AddCredentialsResponse{
		Id: creds.ID,
	}

	return resp, nil
}
