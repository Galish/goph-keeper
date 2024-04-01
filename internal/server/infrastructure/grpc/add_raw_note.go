package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/pkg/logger"

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

	err := s.keeper.AddRawNote(ctx, note)
	if err != nil {
		logger.WithError(err).Error("unable to add binary note")
	}

	if errors.Is(err, keeper.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Id = note.ID

	return &response, nil
}
