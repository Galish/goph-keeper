package keeper

import (
	"context"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"

	"github.com/Galish/goph-keeper/internal/client/entity"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (uc *KeeperUseCase) AddCard(card *entity.Card) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
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

	if _, err := uc.client.AddCard(ctx, req); err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) UpdateCard(card *entity.Card) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
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
	}

	_, err := uc.client.UpdateCard(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) GetCard(id string) (*entity.Card, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetCard(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	card := &entity.Card{
		Title:       resp.Card.GetTitle(),
		Description: resp.Card.GetDescription(),
		Number:      resp.Card.GetNumber(),
		Holder:      resp.Card.GetHolder(),
		CVC:         resp.Card.GetCvc(),
		Expiry:      resp.Card.GetExpiry().AsTime(),
	}

	return card, nil
}

func (uc *KeeperUseCase) GetCardsList() ([]*entity.Card, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
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

func (uc *KeeperUseCase) DeleteCard(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteCard(ctx, req)

	return handleError(err)
}
