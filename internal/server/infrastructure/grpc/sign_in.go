package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
)

func (s *KeeperServer) SignIn(
	ctx context.Context,
	in *pb.AuthRequest,
) (*pb.AuthResponse, error) {
	var response pb.AuthResponse

	credentials := &entity.User{
		Login:    in.Username,
		Password: in.Password,
	}

	token, err := s.user.SignIn(ctx, credentials)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Token = token
	}

	return &response, nil
}
