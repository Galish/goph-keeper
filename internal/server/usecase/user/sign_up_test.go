package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/auth"
)

const secretKey = "secret_key"

var errWriteToRepo = errors.New("failed to write to repo")

func TestSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepository(ctrl)

	m.EXPECT().
		AddUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, user *entity.User) error {
			switch user.Login {
			case "john.doe":
				return nil
			default:
				return errWriteToRepo
			}
		}).
		AnyTimes()

	uc := user.New(m, auth.NewJWTManager(secretKey))

	type want struct {
		err error
	}

	tests := []struct {
		name     string
		username string
		password string
		want     *want
	}{
		{
			"missing credentials",
			"",
			"",
			&want{
				user.ErrMissingCredentials,
			},
		},
		{
			"missing username",
			"",
			"qwe123456",
			&want{
				user.ErrMissingCredentials,
			},
		},
		{
			"missing password",
			"john.doe",
			"",
			&want{
				user.ErrMissingCredentials,
			},
		},
		{
			"valid credentials",
			"john.doe",
			"qwe123456",
			&want{
				nil,
			},
		},
		{
			"write to repo error",
			"johny.doe",
			"qwe123456",
			&want{
				errWriteToRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := uc.SignUp(context.Background(), tt.username, tt.password)

			assert.Equal(t, tt.want.err, err)

			if tt.want.err == nil {
				_, err := uc.Verify(token)

				require.NoError(t, err)
			}
		})
	}
}
