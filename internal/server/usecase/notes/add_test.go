package notes_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
)

var errWriteToRepo = errors.New("failed to write to repo")

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepository(ctrl)

	m.EXPECT().
		Set(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, r *repository.Record) error {
			if r.ID == "#765432" {
				return errWriteToRepo
			}
			return nil
		}).
		AnyTimes()

	uc := notes.New(m)

	tests := []struct {
		name string
		note *entity.Note
		err  error
	}{
		{
			"empty input",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Note{
				Value: "Text note...",
			},
			notes.ErrInvalidEntity,
		},
		{
			"text note",
			&entity.Note{
				ID:    "#12345",
				Title: "Text note",
				Value: "Text note...",
			},
			nil,
		},
		{
			"binary note",
			&entity.Note{
				ID:       "#12345",
				Title:    "Binary note",
				RawValue: []byte("Binary note..."),
			},
			nil,
		},
		{
			"write to repo error",
			&entity.Note{
				ID:    "#765432",
				Title: "Text note",
				Value: "Text note...",
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.Add(context.Background(), tt.note)

			assert.Equal(t, tt.err, err)
		})
	}
}
