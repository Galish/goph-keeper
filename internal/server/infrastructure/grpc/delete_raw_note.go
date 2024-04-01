package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) DeleteRawNote(ctx context.Context, in *pb.DeleteRequest) (*emptypb.Empty, error) {
	user := ctx.Value(interceptors.UserContextKey).(string)

	err := s.keeper.DeleteRawNote(ctx, user, in.GetId())
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"id": in.GetId(),
			}).
			WithError(err).
			Error("unable to delete binary note")
	}

	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return nil, nil
}
