package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) GetCredentialsList(ctx context.Context, _ *emptypb.Empty) (*pb.GetListResponse, error) {
	var response pb.GetListResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	creds, err := s.keeper.GetAllCredentials(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.List = make([]*pb.ListItem, len(creds))

	for i, c := range creds {
		response.List[i] = &pb.ListItem{
			Id:          c.ID,
			Title:       c.Title,
			Description: c.Description,
		}
	}

	return &response, nil
}
