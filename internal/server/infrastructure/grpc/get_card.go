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
)

func (s *KeeperServer) GetCard(ctx context.Context, in *pb.GetRequest) (*pb.GetCardResponse, error) {
	var response pb.GetCardResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	card, err := s.keeper.GetCard(ctx, user, in.Id)
	if errors.Is(err, keeper.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Card = &pb.Card{
		Title:       card.Title,
		Description: card.Description,
		Number:      card.Number,
		Holder:      card.Holder,
		Cvc:         card.CVC,
		Expiry:      timestamppb.New(card.Expiry),
	}

	return &response, nil
}
