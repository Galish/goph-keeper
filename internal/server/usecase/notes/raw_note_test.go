//go:build unit
// +build unit

package notes_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/notes"
)

func TestAddRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		AddSecureNote(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, r *repository.SecureNote) error {
			if r.ID == "#765432" {
				return errWriteToRepo
			}

			return nil
		}).
		AnyTimes()

	uc := notes.New(m)

	tests := []struct {
		name string
		note *entity.RawNote
		err  error
	}{
		{
			"empty input",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				Value: []byte("Hello world!"),
			},
			notes.ErrInvalidEntity,
		},
		{
			"valid entity",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			nil,
		},
		{
			"write to repo error",
			&entity.RawNote{
				ID:    "#765432",
				Title: "Text note",
				Value: []byte("Hello world!"),
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddRawNote(context.Background(), tt.note)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		GetSecureNote(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		DoAndReturn(func(_ context.Context, _, id string, _ repository.SecureNoteType) (*repository.SecureNote, error) {
			switch id {
			case "#23456":
				return &repository.SecureNote{
					ID:      "#23456",
					Type:    repository.TypeRawNote,
					Title:   "Binary note",
					RawNote: []byte("Hello world!"),
				}, nil

			case "#34567":
				return nil, repository.ErrNotFound

			default:
				return nil, errReadFromRepo
			}
		}).
		AnyTimes()

	uc := notes.New(m)

	type want struct {
		note *entity.RawNote
		err  error
	}

	tests := []struct {
		name string
		user string
		id   string
		want *want
	}{
		{
			"missing id",
			"user#12345",
			"",
			&want{
				nil,
				notes.ErrMissingArgument,
			},
		},
		{
			"missing user",
			"",
			"#34567",
			&want{
				nil,
				notes.ErrMissingArgument,
			},
		},
		{
			"nothing found",
			"user#12345",
			"#34567",
			&want{
				nil,
				notes.ErrNotFound,
			},
		},
		{
			"valid binary note",
			"user#12345",
			"#23456",
			&want{
				&entity.RawNote{
					ID:    "#23456",
					Title: "Binary note",
					Value: []byte("Hello world!"),
				},
				nil,
			},
		},
		{
			"read from repo error",
			"user#12345",
			"#435214",
			&want{
				nil,
				errReadFromRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			note, err := uc.GetRawNote(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.want.note, note)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestGetRawNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		GetSecureNotes(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetSecureNotes(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		Return([]*repository.SecureNote{
			{
				ID:      "#12345",
				Type:    repository.TypeRawNote,
				Title:   "Binary note",
				RawNote: []byte("Binary note..."),
			},
			{
				ID:      "#23456",
				Type:    repository.TypeRawNote,
				Title:   "Another binary note",
				RawNote: []byte("Another binary note..."),
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetSecureNotes(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := notes.New(m)

	type want struct {
		notes []*entity.RawNote
		err   error
	}

	tests := []struct {
		name string
		user string
		want *want
	}{
		{
			"missing user",
			"",
			&want{
				nil,
				notes.ErrMissingArgument,
			},
		},
		{
			"empty list",
			"user#12345",
			&want{
				[]*entity.RawNote{},
				nil,
			},
		},
		{
			"list of notes",
			"user#12345",
			&want{
				[]*entity.RawNote{
					{
						ID:    "#12345",
						Title: "Binary note",
						Value: []byte("Binary note..."),
					},
					{
						ID:    "#23456",
						Title: "Another binary note",
						Value: []byte("Another binary note..."),
					},
				},
				nil,
			},
		},
		{
			"read from repo error",
			"user#12345",
			&want{
				nil,
				errReadFromRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			notes, err := uc.GetRawNotes(context.Background(), tt.user)

			assert.Equal(t, tt.want.notes, notes)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestUpdateRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		UpdateSecureNote(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, r *repository.SecureNote) error {
			switch r.ID {
			case "#12345":
				return repository.ErrNotFound

			case "#23456":
				return errWriteToRepo

			case "#34567":
				return repository.ErrVersionConflict

			default:
				return nil
			}
		}).
		AnyTimes()

	uc := notes.New(m)

	tests := []struct {
		name      string
		RawNote   *entity.RawNote
		overwrite bool
		err       error
	}{
		{
			"empty input",
			nil,
			false,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				ID: "#12345",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{
				ID:    "#12345",
				Value: []byte("Hello world!"),
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"nothing found",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			notes.ErrNotFound,
		},
		{
			"version required",
			&entity.RawNote{
				ID:    "#12345",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			false,
			notes.ErrVersionRequired,
		},
		{
			"updated version",
			&entity.RawNote{
				ID:      "#789012",
				Title:   "Binary note",
				Value:   []byte("Hello world!"),
				Version: 10,
			},
			false,
			nil,
		},
		{
			"overwritten",
			&entity.RawNote{
				ID:    "#789012",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			nil,
		},
		{
			"version conflict",
			&entity.RawNote{
				ID:    "#34567",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			notes.ErrVersionConflict,
		},
		{
			"write to repo error",
			&entity.RawNote{
				ID:    "#23456",
				Title: "Binary note",
				Value: []byte("Hello world!"),
			},
			true,
			errWriteToRepo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateRawNote(context.Background(), tt.RawNote, tt.overwrite)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		DeleteSecureNote(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeRawNote)).
		DoAndReturn(func(_ context.Context, _, id string, _ repository.SecureNoteType) error {
			switch id {
			case "#12345":
				return repository.ErrNotFound

			case "#23456":
				return errWriteToRepo

			default:
				return nil
			}
		}).
		AnyTimes()

	uc := notes.New(m)

	tests := []struct {
		name string
		user string
		id   string
		err  error
	}{
		{
			"missing id",
			"user#12345",
			"",
			notes.ErrMissingArgument,
		},
		{
			"missing user",
			"",
			"#12345",
			notes.ErrMissingArgument,
		},
		{
			"nothing found",
			"user#12345",
			"#12345",
			notes.ErrNotFound,
		},
		{
			"write to repo error",
			"user#12345",
			"#23456",
			errWriteToRepo,
		},
		{
			"deleted",
			"user#12345",
			"#34567",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.DeleteRawNote(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.err, err)
		})
	}
}
