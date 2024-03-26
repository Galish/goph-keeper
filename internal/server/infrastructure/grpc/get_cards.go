package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetCards(
	ctx context.Context,
	_ *emptypb.Empty,
) (*pb.GetCardsResponse, error) {
	var response pb.GetCardsResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	cards, err := s.keeper.GetCards(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Cards = make([]*pb.Card, len(cards))

	for i, c := range cards {
		response.Cards[i] = &pb.Card{
			Title:       c.Title,
			Description: c.Description,
			Number:      c.Number,
			Holder:      c.Holder,
			Cvc:         c.CVC,
			Expiry:      timestamppb.New(c.Expiry),
		}
	}

	return &response, nil
}
