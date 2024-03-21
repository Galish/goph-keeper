package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) AddNote(
	ctx context.Context,
	in *pb.AddNoteRequest,
) (*pb.AddNoteResponse, error) {
	var response pb.AddNoteResponse

	note := entity.NewNote()
	note.Title = in.Note.Title
	note.Description = in.Note.Description
	note.Value = in.Note.Value
	note.RawValue = in.Note.RawValue
	note.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	if err := s.keeper.AddNote(ctx, note); err != nil {
		response.Error = err.Error()
	}

	return &response, nil
}
