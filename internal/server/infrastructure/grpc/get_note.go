package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetNote(
	ctx context.Context,
	in *pb.GetNoteRequest,
) (*pb.GetNoteResponse, error) {
	var response pb.GetNoteResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	note, err := s.keeper.GetNote(ctx, user, in.Id)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Note.Title = note.Title
		response.Note.Description = note.Description
		response.Note.Value = note.Value
		response.Note.RawValue = note.RawValue
	}

	return &response, nil
}
