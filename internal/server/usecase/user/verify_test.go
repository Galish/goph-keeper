package user_test

import (
	"errors"
	"testing"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository/mocks"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/auth"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockUserRepository(ctrl)

	uc := user.New(m, auth.NewJWTManager(secretKey))

	type want struct {
		user *entity.User
		err  error
	}

	tests := []struct {
		name  string
		token string
		want  *want
	}{
		{
			"missing token",
			"",
			&want{
				err: errors.New("token contains an invalid number of segments"),
			},
		},
		{
			"invalid token",
			"asdjhd871ije1o9j0assdj",
			&want{
				err: errors.New("token contains an invalid number of segments"),
			},
		},
		{
			"valid token",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIjMTIzNDUifQ.FnPcRyLXm11AqObgLd1HR-OB7FmsPtcbsUg31IUW6Ss",
			&want{
				&entity.User{
					ID: "#12345",
				},
				errors.New("token contains an invalid number of segments"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := uc.Verify(tt.token)

			assert.Equal(t, tt.want.user, user)
			assert.Error(t, tt.want.err, err)
		})
	}
}
