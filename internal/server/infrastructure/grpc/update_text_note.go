package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) UpdateTextNote(ctx context.Context, in *pb.UpdateTextNoteRequest) (*emptypb.Empty, error) {
	note := &entity.TextNote{
		ID:          in.GetId(),
		Title:       in.Note.GetTitle(),
		Description: in.Note.GetDescription(),
		Value:       in.Note.GetValue(),
		CreatedBy:   ctx.Value(interceptors.UserContextKey).(string),
	}

	err := s.keeper.UpdateTextNote(ctx, note)
	if errors.Is(err, keeper.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return nil, nil
}
