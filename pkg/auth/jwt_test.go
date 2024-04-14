//go:build unit
// +build unit

package auth_test

import (
	"errors"
	"testing"

	"github.com/Galish/goph-keeper/pkg/auth"

	"github.com/stretchr/testify/assert"
)

const secretKey = "secret_key"

func TestGenerate(t *testing.T) {
	manager := auth.NewJWTManager(secretKey)

	type want struct {
		token string
		err   error
	}

	tests := []struct {
		name   string
		claims *auth.JWTClaims
		want   *want
	}{
		{
			"missing claims",
			nil,
			&want{
				err: errors.New("missing required argument"),
			},
		},
		{
			"missing user id",
			&auth.JWTClaims{},
			&want{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIifQ.v4mstQ2bHiVOz69p_11a8ZihG1wVUcwgMDnq6ShkseE",
				err:   nil,
			},
		},
		{
			"with user id",
			&auth.JWTClaims{
				UserID: "#12345",
			},
			&want{
				token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIjMTIzNDUifQ.FnPcRyLXm11AqObgLd1HR-OB7FmsPtcbsUg31IUW6Ss",
				err:   nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := manager.Generate(tt.claims)

			assert.Equal(t, tt.want.token, token)
			assert.Equal(t, tt.want.err, err)
		})
	}
}

func TestVerify(t *testing.T) {
	manager := auth.NewJWTManager(secretKey)

	type want struct {
		claims *auth.JWTClaims
		err    error
	}

	tests := []struct {
		name  string
		token string
		want  *want
	}{
		{
			"empty input",
			"",
			&want{
				claims: nil,
				err:    errors.New("token contains an invalid number of segments"),
			},
		},
		{
			"missing user id",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIifQ.v4mstQ2bHiVOz69p_11a8ZihG1wVUcwgMDnq6ShkseE",
			&want{
				claims: &auth.JWTClaims{},
				err:    nil,
			},
		},
		{
			"with user id",
			"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiIjMTIzNDUifQ.FnPcRyLXm11AqObgLd1HR-OB7FmsPtcbsUg31IUW6Ss",
			&want{
				claims: &auth.JWTClaims{
					UserID: "#12345",
				},
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			claims, err := manager.Verify(tt.token)

			assert.Equal(t, tt.want.claims, claims)

			if tt.want.err != nil {
				assert.ErrorContains(t, tt.want.err, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
