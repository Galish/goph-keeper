package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"

	"github.com/Galish/goph-keeper/internal/server/entity"
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
		CreateUser(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, user *entity.User) error {
			switch user.Login {
			case "john.doe":
				return nil
			default:
				return errWriteToRepo
			}
		}).
		AnyTimes()

	uc := user.New(m, secretKey)

	type want struct {
		err error
	}

	tests := []struct {
		name  string
		creds *entity.User
		want  *want
	}{
		{
			"empty input",
			nil,
			&want{
				user.ErrMissingCredentials,
			},
		},
		{
			"missing username",
			&entity.User{
				Login:    "",
				Password: "qwe123456",
			},
			&want{
				user.ErrMissingCredentials,
			},
		},
		{
			"missing password",
			&entity.User{
				Login:    "john.doe",
				Password: "",
			},
			&want{
				user.ErrMissingCredentials,
			},
		},
		{
			"valid credentials",
			&entity.User{
				Login:    "john.doe",
				Password: "qwe123456",
			},
			&want{
				nil,
			},
		},
		{
			"write to repo error",
			&entity.User{
				Login:    "johny.doe",
				Password: "qwe123456",
			},
			&want{
				errWriteToRepo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := uc.SignUp(context.Background(), tt.creds)

			assert.Equal(t, tt.want.err, err)

			if tt.want.err == nil {
				_, err := auth.ParseToken(secretKey, token)

				require.NoError(t, err)
			}
		})
	}
}
