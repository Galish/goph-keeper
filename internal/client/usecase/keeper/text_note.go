package keeper

import (
	"context"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"

	"github.com/Galish/goph-keeper/internal/client/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *KeeperUseCase) AddTextNote(note *entity.TextNote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.AddTextNoteRequest{
		Note: &pb.TextNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
	}

	if _, err := uc.client.AddTextNote(ctx, req); err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) UpdateTextNote(note *entity.TextNote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.UpdateTextNoteRequest{
		Id: note.ID,
		Note: &pb.TextNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
	}

	if _, err := uc.client.UpdateTextNote(ctx, req); err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) GetTextNote(id string) (*entity.TextNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetTextNote(ctx, req)
	if err != nil {
		return nil, err
	}

	note := &entity.TextNote{
		Title:       resp.Note.GetTitle(),
		Description: resp.Note.GetDescription(),
		Value:       resp.Note.GetValue(),
	}

	return note, nil
}

func (uc *KeeperUseCase) DeleteTextNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteTextNote(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (uc *KeeperUseCase) GetTextNotesList() ([]*entity.TextNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	resp, err := uc.client.GetTextNotesList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	var notes = make([]*entity.TextNote, len(resp.GetList()))

	for i, c := range resp.GetList() {
		notes[i] = &entity.TextNote{
			ID:          c.GetId(),
			Title:       c.GetTitle(),
			Description: c.GetDescription(),
		}
	}

	return notes, nil
}
