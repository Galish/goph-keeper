package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddTextNote(ctx context.Context, in *pb.AddTextNoteRequest) (*pb.AddTextNoteResponse, error) {
	note := entity.NewTextNote()
	note.Title = in.Note.GetTitle()
	note.Description = in.Note.GetDescription()
	note.Value = in.Note.GetValue()
	note.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	err := s.keeper.AddTextNote(ctx, note)
	if err != nil {
		logger.WithError(err).Error("unable to add text note")
	}

	if errors.Is(err, keeper.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.AddTextNoteResponse{
		Id: note.ID,
	}

	return resp, nil
}
