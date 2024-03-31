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

func (s *KeeperServer) GetRawNote(ctx context.Context, in *pb.GetRequest) (*pb.GetRawNoteResponse, error) {
	var response pb.GetRawNoteResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	note, err := s.keeper.GetRawNote(ctx, user, in.Id)
	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Note = &pb.RawNote{
		Title:       note.Title,
		Description: note.Description,
		Value:       note.Value,
	}

	return &response, nil
}
