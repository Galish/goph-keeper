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

func TestAddRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		AddRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		AddRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		AddRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		AddRawNote(gomock.Any(), gomock.Any()).
		Return(&pb.AddRawNoteResponse{Id: "#12345"}, nil)

	uc := notes.New(m)

	tests := []struct {
		name    string
		RawNote *entity.RawNote
		err     error
	}{
		{
			"missing entity",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.RawNote{},
			notes.ErrInvalidEntity,
		},
		{
			"failed validation",
			&entity.RawNote{
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			notes.ErrInvalidEntity,
		},
		{
			"writing to repo error",
			&entity.RawNote{
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.RawNote{
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			notes.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.RawNote{
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddRawNote(context.Background(), tt.RawNote)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		UpdateRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		UpdateRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.FailedPrecondition, errors.New("record version conflict").Error()))

	m.EXPECT().
		UpdateRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		UpdateRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		UpdateRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		UpdateRawNote(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	uc := notes.New(m)

	tests := []struct {
		name      string
		RawNote   *entity.RawNote
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
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Credit RawNote",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"missing id",
			&entity.RawNote{
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"failed validation",
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			true,
			notes.ErrInvalidEntity,
		},
		{
			"version conflict",
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			false,
			notes.ErrVersionConflict,
		},
		{
			"nothing found",
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			false,
			notes.ErrNotFound,
		},
		{
			"writing to repo error",
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			true,
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			false,
			notes.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.RawNote{
				ID:    "#12345678",
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			true,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateRawNote(context.Background(), tt.RawNote, tt.overwrite)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		GetRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetRawNote(gomock.Any(), gomock.Any()).
		Return(&pb.GetRawNoteResponse{
			Note: &pb.RawNote{
				Title: "Secret file",
				Value: []byte("Hello world!"),
			},
			Version: 10,
		}, nil)

	uc := notes.New(m)

	type want struct {
		RawNote *entity.RawNote
		err     error
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
				&entity.RawNote{
					Title:   "Secret file",
					Value:   []byte("Hello world!"),
					Version: 10,
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RawNote, err := uc.GetRawNote(context.Background(), tt.id)

			assert.Equal(t, tt.want.RawNote, RawNote)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetRawNotesList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetRawNotesList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{}, nil)

	m.EXPECT().
		GetRawNotesList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetRawNotesList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetRawNotesList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{
			List: []*pb.ListItem{
				{
					Id:          "#12345",
					Title:       "Secret file",
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
		RawNotes []*entity.RawNote
		err      error
	}

	tests := []struct {
		name string
		want *want
	}{
		{
			"nothing found",
			&want{
				[]*entity.RawNote{},
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
				[]*entity.RawNote{
					{
						ID:          "#12345",
						Title:       "Secret file",
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
			RawNotes, err := uc.GetRawNotesList(context.Background())

			assert.Equal(t, tt.want.RawNotes, RawNotes)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteRawNote(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		DeleteRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		DeleteRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		DeleteRawNote(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		DeleteRawNote(gomock.Any(), gomock.Any()).
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
			err := uc.DeleteRawNote(context.Background(), tt.id)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
