package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) SignUp(
	ctx context.Context,
	in *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	token, err := s.user.SignUp(ctx, in.GetUsername(), in.GetPassword())
	if errors.Is(err, user.ErrMissingCredentials) {
		logger.WithError(err).Debug("unable to create the user")
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		logger.WithError(err).Debug("unable to create the user")
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.AccessToken = token

	return &response, nil
}
