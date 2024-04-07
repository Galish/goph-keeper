package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"google.golang.org/grpc"
)

func (c *KeeperClient) SignUp(ctx context.Context, in *pb.AuthRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	resp, err := c.KeeperClient.SignUp(ctx, in)
	if err != nil {
		return nil, err
	}

	c.auth.SetToken(resp.AccessToken)

	return resp, nil
}

func (c *KeeperClient) SignIn(ctx context.Context, in *pb.AuthRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	resp, err := c.KeeperClient.SignIn(ctx, in)
	if err != nil {
		return nil, err
	}

	c.auth.SetToken(resp.AccessToken)

	return resp, nil
}
