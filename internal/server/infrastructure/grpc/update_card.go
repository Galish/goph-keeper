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
	}

	err := s.keeper.UpdateCard(ctx, card)
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"id": in.GetId(),
			}).
			WithError(err).
			Error("unable to update card details")
	}

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
