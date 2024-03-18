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

var errReadFromRepo = errors.New("failed to read from repo")

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepository(ctrl)

	m.EXPECT().
		Get(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, id string) (*repository.Record, error) {
			switch id {
			case "#12345":
				return &repository.Record{
					ID:       "#12345",
					Type:     repository.TypeNote,
					Title:    "Text note",
					TextNote: "Text note...",
				}, nil

			case "#23456":
				return &repository.Record{
					ID:      "#23456",
					Type:    repository.TypeNote,
					Title:   "Binary note",
					RawNote: []byte("Binary note..."),
				}, nil

			case "#34567":
				return &repository.Record{}, nil

			default:
				return nil, errReadFromRepo
			}
		}).
		AnyTimes()

	uc := notes.New(m)

	type want struct {
		note *entity.Note
		err  error
	}

	tests := []struct {
		name string
		id   string
		want *want
	}{
		{
			"invalid type error",
			"#34567",
			&want{
				nil,
				notes.ErrInvalidType,
			},
		},
		{
			"text note",
			"#12345",
			&want{
				&entity.Note{
					ID:    "#12345",
					Title: "Text note",
					Value: "Text note...",
				},
				nil,
			},
		},
		{
			"binary note",
			"#23456",
			&want{
				&entity.Note{
					ID:       "#23456",
					Title:    "Binary note",
					RawValue: []byte("Binary note..."),
				},
				nil,
			},
		},
		{
			"read from repo error",
			"#435214",
			&want{
				nil,
				errReadFromRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := uc.Get(context.Background(), tt.id)

			assert.DeepEqual(t, tt.want.note, note)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
