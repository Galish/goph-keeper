package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddRawNote(ctx context.Context, in *pb.AddRawNoteRequest) (*pb.AddRawNoteResponse, error) {
	var response pb.AddRawNoteResponse

	note := entity.NewRawNote()
	note.Title = in.Note.GetTitle()
	note.Description = in.Note.GetDescription()
	note.Value = in.Note.GetValue()
	note.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	if err := s.keeper.AddRawNote(ctx, note); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Id = note.ID

	return &response, nil
}
