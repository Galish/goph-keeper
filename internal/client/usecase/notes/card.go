package notes

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Galish/goph-keeper/internal/entity"
)

func (uc *KeeperUseCase) AddCard(ctx context.Context, card *entity.Card) error {
	if card == nil || !card.IsValid() {
		return ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.AddCardRequest{
		Card: &pb.Card{
			Title:       card.Title,
			Description: card.Description,
			Number:      card.Number,
			Holder:      card.Holder,
			Cvc:         card.CVC,
			Expiry:      timestamppb.New(card.Expiry),
		},
	}

	_, err := uc.client.AddCard(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) UpdateCard(ctx context.Context, card *entity.Card, overwrite bool) error {
	if card == nil || card.ID == "" || !card.IsValid() {
		return ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.UpdateCardRequest{
		Id: card.ID,
		Card: &pb.Card{
			Title:       card.Title,
			Description: card.Description,
			Number:      card.Number,
			Holder:      card.Holder,
			Cvc:         card.CVC,
			Expiry:      timestamppb.New(card.Expiry),
		},
		Version:   card.Version,
		Overwrite: overwrite,
	}

	_, err := uc.client.UpdateCard(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) GetCard(ctx context.Context, id string) (*entity.Card, error) {
	if id == "" {
		return nil, ErrMissingArgument
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetCard(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	card := &entity.Card{
		Title:       resp.GetCard().GetTitle(),
		Description: resp.GetCard().GetDescription(),
		Number:      resp.GetCard().GetNumber(),
		Holder:      resp.GetCard().GetHolder(),
		CVC:         resp.GetCard().GetCvc(),
		Expiry:      resp.GetCard().GetExpiry().AsTime(),
		Version:     resp.GetVersion(),
	}

	return card, nil
}

func (uc *KeeperUseCase) GetCardsList(ctx context.Context) ([]*entity.Card, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	resp, err := uc.client.GetCardsList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, handleError(err)
	}

	var cards = make([]*entity.Card, len(resp.GetList()))

	for i, c := range resp.GetList() {
		cards[i] = &entity.Card{
			ID:          c.GetId(),
			Title:       c.GetTitle(),
			Description: c.GetDescription(),
		}
	}

	return cards, nil
}

func (uc *KeeperUseCase) DeleteCard(ctx context.Context, id string) error {
	if id == "" {
		return ErrMissingArgument
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteCard(ctx, req)

	return handleError(err)
}
