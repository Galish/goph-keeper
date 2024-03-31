package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) GetTextNote(ctx context.Context, in *pb.GetRequest) (*pb.GetTextNoteResponse, error) {
	var response pb.GetTextNoteResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	note, err := s.keeper.GetTextNote(ctx, user, in.Id)
	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Note = &pb.TextNote{
		Title:       note.Title,
		Description: note.Description,
		Value:       note.Value,
	}

	return &response, nil
}
