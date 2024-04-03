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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *KeeperServer) UpdateCard(ctx context.Context, in *pb.UpdateCardRequest) (*emptypb.Empty, error) {
	card := &entity.Card{
		ID:          in.GetId(),
		Title:       in.Card.GetTitle(),
		Description: in.Card.GetDescription(),
		Number:      in.Card.GetNumber(),
		Holder:      in.Card.GetHolder(),
		CVC:         in.Card.GetCvc(),
		Expiry:      in.Card.GetExpiry().AsTime(),
		CreatedBy:   ctx.Value(interceptors.UserContextKey).(string),
		Version:     in.GetVersion(),
	}

	err := s.keeper.UpdateCard(ctx, card, in.GetOverwrite())
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"id": in.GetId(),
			}).
			WithError(err).
			Error("unable to update card details")
	}

	if errors.Is(err, keeper.ErrInvalidEntity) ||
		errors.Is(err, keeper.ErrVersionRequired) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if errors.Is(err, keeper.ErrVersionConflict) {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return nil, nil
}
