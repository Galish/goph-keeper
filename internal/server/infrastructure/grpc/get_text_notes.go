package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetTextNotes(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.GetNotesResponse, error) {
	var response pb.GetNotesResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	notes, err := s.keeper.GetTextNotes(ctx, user)
	if err != nil {
		response.Error = err.Error()

		return &response, nil
	}

	response.Notes = make([]*pb.Note, len(notes))

	for i, n := range notes {
		response.Notes[i] = &pb.Note{
			Title:       n.Title,
			Description: n.Description,
			Value:       n.Value,
		}
	}

	return &response, nil
}
