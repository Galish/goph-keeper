package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddCredentials(
	ctx context.Context,
	in *pb.AddCredentialsRequest,
) (*pb.AddCredentialsResponse, error) {
	var response pb.AddCredentialsResponse

	creds := entity.NewCredentials()
	creds.Title = in.Credentials.Title
	creds.Description = in.Credentials.Description
	creds.Username = in.Credentials.Username
	creds.Password = in.Credentials.Password
	creds.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	if err := s.keeper.AddCredentials(ctx, creds); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Id = creds.ID

	return &response, nil
}
