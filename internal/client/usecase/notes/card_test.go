package notes_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/Galish/goph-keeper/api/proto"
	mocks "github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/mock"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/entity"
)

var (
	errWriteToRepo  = errors.New("failed to write to repo")
	errReadFromRepo = errors.New("failed to read from repo")
)

func TestAddCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		AddCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

	m.EXPECT().
		AddCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		AddCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		AddCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		AddCard(gomock.Any(), gomock.Any()).
		Return(&pb.AddCardResponse{Id: "#12345"}, nil)

	uc := notes.New(m)

	tests := []struct {
		name string
		card *entity.Card
		err  error
	}{
		{
			"missing entity",
			nil,
			notes.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{},
			notes.ErrInvalidEntity,
		},
		{
			"authorization required",
			&entity.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			notes.ErrAuthRequired,
		},
		{
			"failed validation",
			&entity.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			notes.ErrInvalidEntity,
		},
		{
			"writing to repo error",
			&entity.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			notes.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.AddCard(context.Background(), tt.card)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.FailedPrecondition, errors.New("entity version conflict").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errors.New("failed to write to repo").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, nil)

	uc := notes.New(m)

	tests := []struct {
		name      string
		card      *entity.Card
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
			&entity.Card{
				ID:    "#12345678",
				Title: "Credit card",
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"missing id",
			&entity.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			notes.ErrInvalidEntity,
		},
		{
			"authorization required",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			notes.ErrAuthRequired,
		},
		{
			"failed validation",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			notes.ErrInvalidEntity,
		},
		{
			"version conflict",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			notes.ErrVersionConflict,
		},
		{
			"nothing found",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			notes.ErrNotFound,
		},
		{
			"writing to repo error",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			errWriteToRepo,
		},
		{
			"no internet connection",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			false,
			notes.ErrNoConnection,
		},
		{
			"valid entity",
			&entity.Card{
				ID:     "#12345678",
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				CVC:    "123",
				Expiry: time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
			},
			true,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := uc.UpdateCard(context.Background(), tt.card, tt.overwrite)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

	m.EXPECT().
		GetCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		GetCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetCard(gomock.Any(), gomock.Any()).
		Return(&pb.GetCardResponse{
			Card: &pb.Card{
				Title:  "Credit card",
				Number: "1234 5678 9012 4453",
				Holder: "John Daw",
				Cvc:    "123",
				Expiry: timestamppb.New(time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC)),
			},
			Version: 10,
		}, nil)

	uc := notes.New(m)

	type want struct {
		card *entity.Card
		err  error
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
			"authorization required",
			"#12345",
			&want{
				nil,
				notes.ErrAuthRequired,
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
				&entity.Card{
					Title:   "Credit card",
					Number:  "1234 5678 9012 4453",
					Holder:  "John Daw",
					CVC:     "123",
					Expiry:  time.Date(2025, time.Month(6), 11, 0, 0, 0, 0, time.UTC),
					Version: 10,
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card, err := uc.GetCard(context.Background(), tt.id)

			assert.Equal(t, tt.want.card, card)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetCardsList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		GetCardsList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

	m.EXPECT().
		GetCardsList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{}, nil)

	m.EXPECT().
		GetCardsList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		GetCardsList(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		GetCardsList(gomock.Any(), gomock.Any()).
		Return(&pb.GetListResponse{
			List: []*pb.ListItem{
				{
					Id:          "#12345",
					Title:       "Sberbank",
					Description: "Credit card",
				},
				{
					Id:          "#23456",
					Title:       "Tinkoff",
					Description: "Debit card",
				},
			},
		}, nil)

	uc := notes.New(m)

	type want struct {
		cards []*entity.Card
		err   error
	}

	tests := []struct {
		name string
		want *want
	}{
		{
			"authorization required",
			&want{
				nil,
				notes.ErrAuthRequired,
			},
		},
		{
			"nothing found",
			&want{
				[]*entity.Card{},
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
				[]*entity.Card{
					{
						ID:          "#12345",
						Title:       "Sberbank",
						Description: "Credit card",
					},
					{
						ID:          "#23456",
						Title:       "Tinkoff",
						Description: "Debit card",
					},
				},
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cards, err := uc.GetCardsList(context.Background())

			assert.Equal(t, tt.want.cards, cards)

			if err != nil {
				assert.Equal(t, err, tt.want.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestDeleteCard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockKeeperClient(ctrl)

	m.EXPECT().
		DeleteCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unauthenticated, errors.New("authorization required").Error()))

	m.EXPECT().
		DeleteCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.NotFound, errors.New("no entity found").Error()))

	m.EXPECT().
		DeleteCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Internal, errReadFromRepo.Error()))

	m.EXPECT().
		DeleteCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.Unavailable, errors.New("no connection").Error()))

	m.EXPECT().
		DeleteCard(gomock.Any(), gomock.Any()).
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
			"authorization required",
			"#12345",
			notes.ErrAuthRequired,
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
			err := uc.DeleteCard(context.Background(), tt.id)

			if err != nil {
				assert.Equal(t, err, tt.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
