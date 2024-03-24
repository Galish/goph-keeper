package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) GetNote(
	ctx context.Context,
	in *pb.GetNoteRequest,
) (*pb.GetNoteResponse, error) {
	var response pb.GetNoteResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	note, err := s.keeper.GetNote(ctx, user, in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Note = &pb.Note{
		Title:       note.Title,
		Description: note.Description,
		Value:       note.Value,
		RawValue:    note.RawValue,
	}

	return &response, nil
}
