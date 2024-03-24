package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetRawNotes(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.GetNotesResponse, error) {
	var response pb.GetNotesResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	notes, err := s.keeper.GetRawNotes(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Notes = make([]*pb.Note, len(notes))

	for i, n := range notes {
		response.Notes[i] = &pb.Note{
			Title:       n.Title,
			Description: n.Description,
			RawValue:    n.RawValue,
		}
	}

	return &response, nil
}
