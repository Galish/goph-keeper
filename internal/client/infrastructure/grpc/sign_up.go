package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"
)

func (c *KeeperClient) SignUp(username, password string) (string, error) {
	resp, err := c.service.SignUp(
		context.Background(),
		&pb.AuthRequest{
			Username: username,
			Password: password,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.AccessToken, nil
}
