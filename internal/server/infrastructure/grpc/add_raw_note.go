package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddRawNote(ctx context.Context, in *pb.AddRawNoteRequest) (*pb.AddRawNoteResponse, error) {
	note := entity.NewRawNote()
	note.Title = in.GetNote().GetTitle()
	note.Description = in.GetNote().GetDescription()
	note.Value = in.GetNote().GetValue()
	note.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	err := s.notes.AddRawNote(ctx, note)
	if err != nil {
		logger.WithError(err).Error("unable to add binary note")
	}

	if errors.Is(err, notes.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.AddRawNoteResponse{
		Id: note.ID,
	}

	return resp, nil
}
