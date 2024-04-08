package keeper

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Galish/goph-keeper/internal/entity"
)

func (uc *KeeperUseCase) AddTextNote(ctx context.Context, note *entity.TextNote) error {
	if note == nil || !note.IsValid() {
		return ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

func (uc *KeeperUseCase) UpdateTextNote(ctx context.Context, note *entity.TextNote, overwrite bool) error {
	if note == nil || note.ID == "" || !note.IsValid() {
		return ErrInvalidEntity
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

func (uc *KeeperUseCase) GetTextNote(ctx context.Context, id string) (*entity.TextNote, error) {
	if id == "" {
		return nil, ErrMissingArgument
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.GetRequest{
		Id: id,
	}

	resp, err := uc.client.GetTextNote(ctx, req)
	if err != nil {
		return nil, handleError(err)
	}

	note := &entity.TextNote{
		Title:       resp.GetNote().GetTitle(),
		Description: resp.GetNote().GetDescription(),
		Value:       resp.GetNote().GetValue(),
		Version:     resp.GetVersion(),
	}

	return note, nil
}

func (uc *KeeperUseCase) GetTextNotesList(ctx context.Context) ([]*entity.TextNote, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

func (uc *KeeperUseCase) DeleteTextNote(ctx context.Context, id string) error {
	if id == "" {
		return ErrMissingArgument
	}

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteTextNote(ctx, req)

	return handleError(err)
}
