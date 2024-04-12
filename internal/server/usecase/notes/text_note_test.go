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

func TestAddTextNote(t *testing.T) {
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
		note *entity.TextNote
		err  error
	}{
		{
			"empty input",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.TextNote{
				Value: "Text note...",
			},
			notes.ErrInvalidEntity,
		},
		{
			"valid entity",
			&entity.TextNote{
				ID:    "#12345",
				Title: "Text note",
				Value: "Text note...",
			},
			nil,
		},
		{
			"write to repo error",
			&entity.TextNote{
				ID:    "#765432",
				Title: "Text note",
				Value: "Text note...",
			},
			errWriteToRepo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddTextNote(context.Background(), tt.note)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestGetTextNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		GetSecureNote(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeTextNote)).
		DoAndReturn(func(_ context.Context, _, id string, _ repository.SecureNoteType) (*repository.SecureNote, error) {
			switch id {
			case "#12345":
				return &repository.SecureNote{
					ID:       "#12345",
					Type:     repository.TypeTextNote,
					Title:    "Text note",
					TextNote: "Text note...",
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
		note *entity.TextNote
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
			"valid text note",
			"user#12345",
			"#12345",
			&want{
				&entity.TextNote{
					ID:    "#12345",
					Title: "Text note",
					Value: "Text note...",
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
			note, err := uc.GetTextNote(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.want.note, note)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestGetTextNotes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		GetSecureNotes(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeTextNote)).
		Return(nil, nil).
		Times(1)

	m.EXPECT().
		GetSecureNotes(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeTextNote)).
		Return([]*repository.SecureNote{
			{
				ID:       "#12345",
				Type:     repository.TypeTextNote,
				Title:    "Text note",
				TextNote: "Text note...",
			},
			{
				ID:       "#23456",
				Type:     repository.TypeTextNote,
				Title:    "Another text note",
				TextNote: "Another text note...",
			},
		}, nil).
		Times(1)

	m.EXPECT().
		GetSecureNotes(gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeTextNote)).
		Return(nil, errReadFromRepo).
		Times(1)

	uc := notes.New(m)

	type want struct {
		notes []*entity.TextNote
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
				[]*entity.TextNote{},
				nil,
			},
		},
		{
			"list of notes",
			"user#12345",
			&want{
				[]*entity.TextNote{
					{
						ID:    "#12345",
						Title: "Text note",
						Value: "Text note...",
					},
					{
						ID:    "#23456",
						Title: "Another text note",
						Value: "Another text note...",
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
			notes, err := uc.GetTextNotes(context.Background(), tt.user)

			assert.Equal(t, tt.want.notes, notes)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestUpdateTextNote(t *testing.T) {
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
		TextNote  *entity.TextNote
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
			&entity.TextNote{
				ID: "#12345",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.TextNote{
				ID:    "#12345",
				Title: "Text note",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.TextNote{
				ID:    "#12345",
				Value: "Text note...",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"nothing found",
			&entity.TextNote{
				ID:    "#12345",
				Title: "Text note",
				Value: "Text note...",
			},
			true,
			notes.ErrNotFound,
		},
		{
			"version required",
			&entity.TextNote{
				ID:    "#12345",
				Title: "Text note",
				Value: "Text note...",
			},
			false,
			notes.ErrVersionRequired,
		},
		{
			"updated version",
			&entity.TextNote{
				ID:      "#789012",
				Title:   "Text note",
				Value:   "Text note...",
				Version: 10,
			},
			false,
			nil,
		},
		{
			"overwritten",
			&entity.TextNote{
				ID:    "#789012",
				Title: "Text note",
				Value: "Text note...",
			},
			true,
			nil,
		},
		{
			"version conflict",
			&entity.TextNote{
				ID:    "#34567",
				Title: "Text note",
				Value: "Text note...",
			},
			true,
			notes.ErrVersionConflict,
		},
		{
			"write to repo error",
			&entity.TextNote{
				ID:    "#23456",
				Title: "Text note",
				Value: "Text note...",
			},
			true,
			errWriteToRepo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateTextNote(context.Background(), tt.TextNote, tt.overwrite)

			assert.Equal(t, tt.err, err)
		})
	}
}

func TestDeleteTextNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockSecureNotesRepository(ctrl)

	m.EXPECT().
		DeleteSecureNote(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Eq(repository.TypeTextNote)).
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
			err := uc.DeleteTextNote(context.Background(), tt.user, tt.id)

			assert.Equal(t, tt.err, err)
		})
	}
}
