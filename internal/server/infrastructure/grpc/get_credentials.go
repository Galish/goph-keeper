package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) GetCredentials(
	ctx context.Context,
	in *pb.GetCredentialsRequest,
) (*pb.GetCredentialsResponse, error) {
	var response pb.GetCredentialsResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	creds, err := s.keeper.GetCredentials(ctx, user, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Credentials = &pb.Credentials{
		Title:       creds.Title,
		Description: creds.Description,
		Username:    creds.Username,
		Password:    creds.Password,
	}

	return &response, nil
}
