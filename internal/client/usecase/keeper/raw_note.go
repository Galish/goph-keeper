package keeper

import (
	"context"
	"time"

	pb "github.com/Galish/goph-keeper/api/proto"

	"github.com/Galish/goph-keeper/internal/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *KeeperUseCase) AddRawNote(note *entity.RawNote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.AddRawNoteRequest{
		Note: &pb.RawNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
	}

	_, err := uc.client.AddRawNote(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) UpdateRawNote(note *entity.RawNote) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.UpdateRawNoteRequest{
		Id: note.ID,
		Note: &pb.RawNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
	}

	_, err := uc.client.UpdateRawNote(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) GetRawNote(id string) (*entity.RawNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetRawNote(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	note := &entity.RawNote{
		Title:       resp.Note.GetTitle(),
		Description: resp.Note.GetDescription(),
		Value:       resp.Note.GetValue(),
	}

	return note, nil
}

func (uc *KeeperUseCase) GetRawNotesList() ([]*entity.RawNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	resp, err := uc.client.GetRawNotesList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, handleError(err)
	}

	var notes = make([]*entity.RawNote, len(resp.GetList()))

	for i, c := range resp.GetList() {
		notes[i] = &entity.RawNote{
			ID:          c.GetId(),
			Title:       c.GetTitle(),
			Description: c.GetDescription(),
		}
	}

	return notes, nil
}

func (uc *KeeperUseCase) DeleteRawNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteRawNote(ctx, req)

	return handleError(err)
}
