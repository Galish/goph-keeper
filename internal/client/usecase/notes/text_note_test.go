package notes_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Galish/goph-keeper/api/proto"
	mocks "github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/mock"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/entity"
)

func TestAddTextNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		AddTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		AddTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		AddTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		AddTextNote(gomock.Any(), gomock.Any()).
		Return(&pb.AddTextNoteResponse{Id: "#12345"}, nil)

	uc := notes.New(m)

	tests := []struct {
		name     string
		TextNote *entity.TextNote
		err      error
	}{
		{
			"missing entity",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.TextNote{},
			notes.ErrInvalidEntity,
		},
		{
			"failed validation",
			&entity.TextNote{
				Title: "My text note",
				Value: "Text note ...",
			},
			notes.ErrInvalidEntity,
		},
		{
			"writing to repo error",
			&entity.TextNote{
				Title: "My text note",
				Value: "Text note ...",
			},
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.TextNote{
				Title: "My text note",
				Value: "Text note ...",
			},
			notes.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.TextNote{
				Title: "My text note",
				Value: "Text note ...",
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddTextNote(context.Background(), tt.TextNote)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateTextNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		UpdateTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		UpdateTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.FailedPrecondition, errors.New("record version conflict").Error()))

	m.EXPECT().
		UpdateTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		UpdateTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		UpdateTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		UpdateTextNote(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	uc := notes.New(m)

	tests := []struct {
		name      string
		TextNote  *entity.TextNote
		overwrite bool
		err       error
	}{
		{
			"missing entity",
			nil,
			false,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "Credit TextNote",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"missing id",
			&entity.TextNote{
				Title: "My text note",
				Value: "Text note ...",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"failed validation",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "My text note",
				Value: "Text note ...",
			},
			true,
			notes.ErrInvalidEntity,
		},
		{
			"version conflict",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "My text note",
				Value: "Text note ...",
			},
			false,
			notes.ErrVersionConflict,
		},
		{
			"nothing found",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "My text note",
				Value: "Text note ...",
			},
			false,
			notes.ErrNotFound,
		},
		{
			"writing to repo error",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "My text note",
				Value: "Text note ...",
			},
			true,
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "My text note",
				Value: "Text note ...",
			},
			false,
			notes.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.TextNote{
				ID:    "#12345678",
				Title: "My text note",
				Value: "Text note ...",
			},
			true,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateTextNote(context.Background(), tt.TextNote, tt.overwrite)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTextNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		GetTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetTextNote(gomock.Any(), gomock.Any()).
		Return(&pb.GetTextNoteResponse{
			Note: &pb.TextNote{
				Title: "My text note",
				Value: "Text note ...",
			},
			Version: 10,
		}, nil)

	uc := notes.New(m)

	type want struct {
		TextNote *entity.TextNote
		err      error
	}

	tests := []struct {
		name string
		id   string
		want *want
	}{
		{
			"missing argument",
			"",
			&want{
				nil,
				notes.ErrMissingArgument,
			},
		},
		{
			"nothing found",
			"#12345",
			&want{
				nil,
				notes.ErrNotFound,
			},
		},
		{
			"reading from repo error",
			"#12345",
			&want{
				nil,
				errReadFromRepo,
			},
		},
		{
			"no internet connection",
			"#12345",
			&want{
				nil,
				notes.ErrNoConnection,
			},
		},
		{
			"valid entity",
			"#12345",
			&want{
				&entity.TextNote{
					Title:   "My text note",
					Value:   "Text note ...",
					Version: 10,
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TextNote, err := uc.GetTextNote(context.Background(), tt.id)

			assert.Equal(t, tt.want.TextNote, TextNote)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetTextNotesList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetTextNotesList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{}, nil)

	m.EXPECT().
		GetTextNotesList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetTextNotesList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetTextNotesList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{
			List: []*pb.ListItem{
				{
					Id:          "#12345",
					Title:       "My text note",
					Description: "",
				},
				{
					Id:          "#23456",
					Title:       "Another file",
					Description: "Super secret file",
				},
			},
		}, nil)

	uc := notes.New(m)

	type want struct {
		TextNotes []*entity.TextNote
		err       error
	}

	tests := []struct {
		name string
		want *want
	}{
		{
			"nothing found",
			&want{
				[]*entity.TextNote{},
				nil,
			},
		},
		{
			"reading from repo error",
			&want{
				nil,
				errReadFromRepo,
			},
		},
		{
			"no internet connection",
			&want{
				nil,
				notes.ErrNoConnection,
			},
		},
		{
			"valid entity",
			&want{
				[]*entity.TextNote{
					{
						ID:          "#12345",
						Title:       "My text note",
						Description: "",
					},
					{
						ID:          "#23456",
						Title:       "Another file",
						Description: "Super secret file",
					},
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TextNotes, err := uc.GetTextNotesList(context.Background())

			assert.Equal(t, tt.want.TextNotes, TextNotes)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteTextNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		DeleteTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		DeleteTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		DeleteTextNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		DeleteTextNote(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	uc := notes.New(m)

	tests := []struct {
		name string
		id   string
		err  error
	}{
		{
			"missing argument",
			"",
			notes.ErrMissingArgument,
		},
		{
			"nothing found",
			"#12345",
			notes.ErrNotFound,
		},
		{
			"reading from repo error",
			"#12345",
			errReadFromRepo,
		},
		{
			"no internet connection",
			"#12345",
			notes.ErrNoConnection,
		},
		{
			"successfully deleted",
			"#12345",
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.DeleteTextNote(context.Background(), tt.id)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
