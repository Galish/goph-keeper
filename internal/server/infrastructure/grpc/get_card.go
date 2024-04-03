package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/pkg/logger"
)

func (s *KeeperServer) GetCard(ctx context.Context, in *pb.GetRequest) (*pb.GetCardResponse, error) {
	user := ctx.Value(interceptors.UserContextKey).(string)

	card, err := s.keeper.GetCard(ctx, user, in.GetId())
	if err != nil {
		logger.
			WithFields(logger.Fields{
				"id": in.GetId(),
			}).
			WithError(err).
			Error("unable to get card")
	}

	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := pb.GetCardResponse{
		Card: &pb.Card{
			Title:       card.Title,
			Description: card.Description,
			Number:      card.Number,
			Holder:      card.Holder,
			Cvc:         card.CVC,
			Expiry:      timestamppb.New(card.Expiry),
		},
		Version: card.Version,
	}

	return &resp, nil
}
