package keeper

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"github.com/Galish/goph-keeper/internal/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (uc *KeeperUseCase) AddRawNote(ctx context.Context, note *entity.RawNote) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

func (uc *KeeperUseCase) UpdateRawNote(ctx context.Context, note *entity.RawNote, overwrite bool) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.UpdateRawNoteRequest{
		Id: note.ID,
		Note: &pb.RawNote{
			Title:       note.Title,
			Description: note.Description,
			Value:       note.Value,
		},
		Version:   note.Version,
		Overwrite: overwrite,
	}

	_, err := uc.client.UpdateRawNote(ctx, req)

	return handleError(err)
}

func (uc *KeeperUseCase) GetRawNote(ctx context.Context, id string) (*entity.RawNote, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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
		Version:     resp.GetVersion(),
	}

	return note, nil
}

func (uc *KeeperUseCase) GetRawNotesList(ctx context.Context) ([]*entity.RawNote, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
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

func (uc *KeeperUseCase) DeleteRawNote(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	req := &pb.DeleteRequest{
		Id: id,
	}

	_, err := uc.client.DeleteRawNote(ctx, req)

	return handleError(err)
}
