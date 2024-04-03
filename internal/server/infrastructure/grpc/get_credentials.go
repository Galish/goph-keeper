package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) GetCredentials(ctx context.Context, in *pb.GetRequest) (*pb.GetCredentialsResponse, error) {
	user := ctx.Value(interceptors.UserContextKey).(string)

	creds, err := s.keeper.GetCredentials(ctx, user, in.GetId())
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"id": in.GetId(),
			}).
			WithError(err).
			Error("unable to get credentials")
	}

	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := pb.GetCredentialsResponse{
		Credentials: &pb.Credentials{
			Title:       creds.Title,
			Description: creds.Description,
			Username:    creds.Username,
			Password:    creds.Password,
		},
		Version: creds.Version,
	}

	return &resp, nil
}
