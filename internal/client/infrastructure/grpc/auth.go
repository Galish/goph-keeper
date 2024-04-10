package grpc

import (
	"context"

	pb "github.com/Galish/goph-keeper/api/proto"

	"google.golang.org/grpc"
)

// SignUp implements the gRPC client registration method.
func (c *KeeperClient) SignUp(ctx context.Context, in *pb.AuthRequest, _ ...grpc.CallOption) (*pb.AuthResponse, error) {
	resp, err := c.KeeperClient.SignUp(ctx, in)
	if err != nil {
		return nil, err
	}

	c.auth.SetToken(resp.GetAccessToken())

	return resp, nil
}

// SignIn implements the gRPC client authorization method.
func (c *KeeperClient) SignIn(ctx context.Context, in *pb.AuthRequest, _ ...grpc.CallOption) (*pb.AuthResponse, error) {
	resp, err := c.KeeperClient.SignIn(ctx, in)
	if err != nil {
		return nil, err
	}

	c.auth.SetToken(resp.GetAccessToken())

	return resp, nil
}
