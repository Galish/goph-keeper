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

func (s *KeeperServer) AddCard(ctx context.Context, in *pb.AddCardRequest) (*pb.AddCardResponse, error) {
	card := entity.NewCard()
	card.Title = in.Card.GetTitle()
	card.Description = in.Card.GetDescription()
	card.Number = in.Card.GetNumber()
	card.Holder = in.Card.GetHolder()
	card.CVC = in.Card.GetCvc()
	card.Expiry = in.Card.Expiry.AsTime()
	card.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	err := s.keeper.AddCard(ctx, card)
	if err != nil {
		logger.WithError(err).Error("unable to add card")
	}

	if errors.Is(err, keeper.ErrInvalidEntity) {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.AddCardResponse{
		Id: card.ID,
	}

	return resp, nil
}
