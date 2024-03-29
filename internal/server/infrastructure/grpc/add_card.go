package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddCard(
	ctx context.Context,
	in *pb.AddCardRequest,
) (*pb.AddCardResponse, error) {
	var response pb.AddCardResponse

	card := entity.NewCard()
	card.Title = in.Card.Title
	card.Description = in.Card.Description
	card.Number = in.Card.Number
	card.Holder = in.Card.Holder
	card.CVC = in.Card.Cvc
	card.Expiry = in.Card.Expiry.AsTime()
	card.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	if err := s.keeper.AddCard(ctx, card); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	response.Id = card.ID

	return &response, nil
}
