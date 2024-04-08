package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) GetRawNote(ctx context.Context, in *pb.GetRequest) (*pb.GetRawNoteResponse, error) {
	user := ctx.Value(interceptors.UserContextKey).(string)

	note, err := s.notes.GetRawNote(ctx, user, in.Id)
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"id": in.GetId(),
			}).
			WithError(err).
			Error("unable to get binary note")
	}

	if errors.Is(err, notes.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := pb.GetRawNoteResponse{
		Note: &pb.RawNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
		Version: note.Version,
	}

	return &resp, nil
}
