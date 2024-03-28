package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) GetAllCredentials(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.GetAllCredentialsResponse, error) {
	var response pb.GetAllCredentialsResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	creds, err := s.keeper.GetAllCredentials(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Credentials = make([]*pb.Credentials, len(creds))

	for i, c := range creds {
		response.Credentials[i] = &pb.Credentials{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			Username:    c.Username,
			Password:    c.Password,
		}
	}

	return &response, nil
}
