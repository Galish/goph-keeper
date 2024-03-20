package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) AddCard(
	ctx context.Context,
	in *pb.AddCardRequest,
) (*pb.AddCardResponse, error) {
	var response pb.AddCardResponse

	card := entity.NewCard()
	card.Number = in.Card.Number
	card.Holder = in.Card.Holder
	card.CVC = in.Card.Cvc
	// card.Expiry = in.Card.Expiry
	card.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	if err := s.keeper.AddCard(ctx, card); err != nil {
		response.Error = err.Error()
	}

	return &response, nil
}
