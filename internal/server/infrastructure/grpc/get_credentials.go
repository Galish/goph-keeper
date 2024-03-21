package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetCredentials(
	ctx context.Context,
	in *pb.GetCredentialsRequest,
) (*pb.GetCredentialsResponse, error) {
	var response pb.GetCredentialsResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	creds, err := s.keeper.GetCredentials(ctx, user, in.Id)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Credentials.Title = creds.Title
		response.Credentials.Description = creds.Description
		response.Credentials.Username = creds.Username
		response.Credentials.Password = creds.Password
	}

	return &response, nil
}
