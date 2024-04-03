package keeper

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"github.com/Galish/goph-keeper/internal/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *KeeperUseCase) AddTextNote(note *entity.TextNote) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.AddTextNoteRequest{
		Note: &pb.TextNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
	}

	_, err := uc.client.AddTextNote(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) UpdateTextNote(note *entity.TextNote, overwrite bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.UpdateTextNoteRequest{
		Id: note.ID,
		Note: &pb.TextNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
		Version:   note.Version,
		Overwrite: overwrite,
	}

	_, err := uc.client.UpdateTextNote(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) GetTextNote(id string) (*entity.TextNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetTextNote(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	note := &entity.TextNote{
		Title:       resp.Note.GetTitle(),
		Description: resp.Note.GetDescription(),
		Value:       resp.Note.GetValue(),
		Version:     resp.GetVersion(),
	}

	return note, nil
}

func (uc *KeeperUseCase) GetTextNotesList() ([]*entity.TextNote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	resp, err := uc.client.GetTextNotesList(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, handleError(err)
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

func (uc *KeeperUseCase) DeleteTextNote(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDefault)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteTextNote(ctx, req)

	return handleError(err)
}
