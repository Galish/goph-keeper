package keeper_test

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
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
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

	uc := keeper.New(m)

	tests := []struct {
		name string
		card *entity.Card
		err  error
	}{
		{
			"missing entity",
			nil,
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{},
			keeper.ErrInvalidEntity,
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
			keeper.ErrInvalidEntity,
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
			keeper.ErrNoConnection,
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
		Return(nil, status.Error(codes.InvalidArgument, errors.New("failed entity validation").Error()))

	m.EXPECT().
		UpdateCard(gomock.Any(), gomock.Any()).
		Return(nil, status.Error(codes.FailedPrecondition, errors.New("record version conflict").Error()))

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

	uc := keeper.New(m)

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
			keeper.ErrInvalidEntity,
		},
		{
			"invalid entity",
			&entity.Card{
				ID:    "#12345678",
				Title: "Credit card",
			},
			false,
			keeper.ErrInvalidEntity,
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
			keeper.ErrInvalidEntity,
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
			keeper.ErrInvalidEntity,
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
			keeper.ErrVersionConflict,
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
			keeper.ErrNotFound,
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
			keeper.ErrNoConnection,
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

	uc := keeper.New(m)

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
				keeper.ErrMissingArgument,
			},
		},
		{
			"nothing found",
			"#12345",
			&want{
				nil,
				keeper.ErrNotFound,
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
				keeper.ErrNoConnection,
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

	uc := keeper.New(m)

	type want struct {
		cards []*entity.Card
		err   error
	}

	tests := []struct {
		name string
		want *want
	}{
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
				keeper.ErrNoConnection,
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

	uc := keeper.New(m)

	tests := []struct {
		name string
		id   string
		err  error
	}{
		{
			"missing argument",
			"",
			keeper.ErrMissingArgument,
		},
		{
			"nothing found",
			"#12345",
			keeper.ErrNotFound,
		},
		{
			"reading from repo error",
			"#12345",
			errReadFromRepo,
		},
		{
			"no internet connection",
			"#12345",
			keeper.ErrNoConnection,
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