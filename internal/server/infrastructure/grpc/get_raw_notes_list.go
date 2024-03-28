package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetRawNotesList(ctx context.Context, _ *emptypb.Empty) (*pb.GetListResponse, error) {
	var response pb.GetListResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	notes, err := s.keeper.GetRawNotes(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.List = make([]*pb.ListItem, len(notes))

	for i, n := range notes {
		response.List[i] = &pb.ListItem{
			Id:          n.ID,
			Title:       n.Title,
			Description: n.Description,
		}
	}

	return &response, nil
}
