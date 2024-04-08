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

func (s *KeeperServer) SignUp(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	token, err := s.user.SignUp(ctx, in.GetUsername(), in.GetPassword())
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"username": in.GetUsername(),
				"password": in.GetPassword(),
			}).
			WithError(err).
			Error("unable to sign up")
	}

	if errors.Is(err, user.ErrMissingCredentials) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, user.ErrConflict) {
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.AccessToken = token

	return &response, nil
}
