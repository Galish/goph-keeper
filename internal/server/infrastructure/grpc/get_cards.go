package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

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
		response.Error = err.Error()

		return &response, nil
	}

	response.Cards = make([]*pb.Card, len(cards))

	for i, c := range cards {
		response.Cards[i] = &pb.Card{
			Number: c.Number,
			Holder: c.Holder,
			Cvc:    c.CVC,
			// Expiry: c.Expiry,
		}
	}

	return &response, nil
}
