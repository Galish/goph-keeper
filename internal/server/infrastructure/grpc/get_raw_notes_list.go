package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/pkg/logger"
)

func (s *KeeperServer) GetRawNotesList(ctx context.Context, _ *emptypb.Empty) (*pb.GetListResponse, error) {
	var response pb.GetListResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	notes, err := s.keeper.GetRawNotes(ctx, user)
	if err != nil {
		logger.WithError(err).Error("unable to get binary notes")

		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.List = make([]*pb.ListItem, len(notes))

	for i, n := range notes {
		response.List[i] = &pb.ListItem{
			Id:           n.ID,
			Title:        n.Title,
			Description:  n.Description,
			CreatedAt:    timestamppb.New(n.CreatedAt),
			LastEditedAt: timestamppb.New(n.LastEditedAt),
		}
	}

	return &response, nil
}
