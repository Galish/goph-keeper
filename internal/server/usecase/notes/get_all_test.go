package notes_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
)

func TestGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockRepository(ctrl)

	m.EXPECT().
		GetByType(gomock.Any(), gomock.Eq(repository.TypeNote)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetByType(gomock.Any(), gomock.Eq(repository.TypeNote)).
		Return([]*repository.Record{
			{
				ID:       "#12345",
				Type:     repository.TypeNote,
				Title:    "Text note",
				TextNote: "Text note...",
			},
			{
				ID:      "#23456",
				Type:    repository.TypeNote,
				Title:   "Binary note",
				RawNote: []byte("Binary note..."),
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetByType(gomock.Any(), gomock.Eq(repository.TypeNote)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := notes.New(m)

	type want struct {
		notes []*entity.Note
		err   error
	}

	tests := []struct {
		name string
		want *want
	}{
		{
			"empty list",
			&want{
				[]*entity.Note{},
				nil,
			},
		},
		{
			"list of notes",
			&want{
				[]*entity.Note{
					{
						ID:    "#12345",
						Title: "Text note",
						Value: "Text note...",
					},
					{
						ID:       "#23456",
						Title:    "Binary note",
						RawValue: []byte("Binary note..."),
					},
				},
				nil,
			},
		},
		{
			"read from repo error",
			&want{
				nil,
				errReadFromRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notes, err := uc.GetAll(context.Background())

			assert.DeepEqual(t, tt.want.notes, notes)
			assert.Equal(t, tt.want.err, err)
		})
	}
}
