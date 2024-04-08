package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *KeeperServer) GetCredentialsList(ctx context.Context, _ *emptypb.Empty) (*pb.GetListResponse, error) {
	var response pb.GetListResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	creds, err := s.notes.GetAllCredentials(ctx, user)
	if err != nil {
		logger.WithError(err).Error("unable to get credentials list")

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.List = make([]*pb.ListItem, len(creds))

	for i, c := range creds {
		response.List[i] = &pb.ListItem{
			Id:           c.ID,
			Title:        c.Title,
			Description:  c.Description,
			CreatedAt:    timestamppb.New(c.CreatedAt),
			LastEditedAt: timestamppb.New(c.LastEditedAt),
		}
	}

	return &response, nil
}
