package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
)

func (s *KeeperServer) GetCard(
	ctx context.Context,
	in *pb.GetCardRequest,
) (*pb.GetCardResponse, error) {
	var response pb.GetCardResponse

	user := ctx.Value(interceptors.UserContextKey).(string)

	card, err := s.keeper.GetCard(ctx, user, in.Id)
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Card.Title = card.Title
		response.Card.Description = card.Description
		response.Card.Number = card.Number
		response.Card.Holder = card.Holder
		response.Card.Cvc = card.CVC
		response.Card.Expiry = timestamppb.New(card.Expiry)
	}

	return &response, nil
}
