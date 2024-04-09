package grpc

import (
	"context"
	"errors"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *KeeperServer) AddCard(ctx context.Context, in *pb.AddCardRequest) (*pb.AddCardResponse, error) {
	card := entity.NewCard()
	card.Title = in.GetCard().GetTitle()
	card.Description = in.GetCard().GetDescription()
	card.Number = in.GetCard().GetNumber()
	card.Holder = in.GetCard().GetHolder()
	card.CVC = in.GetCard().GetCvc()
	card.Expiry = in.GetCard().GetExpiry().AsTime()
	card.CreatedBy = ctx.Value(interceptors.UserContextKey).(string)

	err := s.notes.AddCard(ctx, card)
	if err != nil {
		logger.WithError(err).Error("unable to add card")
	}

	if errors.Is(err, notes.ErrInvalidEntity) {
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
