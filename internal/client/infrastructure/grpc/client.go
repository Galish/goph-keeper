package grpc

import (
	"context"
	"log"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/interceptors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type KeeperClient struct {
	pb.KeeperClient
	authClient *auth.AuthClient
	conn       *grpc.ClientConn
}

func NewClient(cfg *config.Config, authClient *auth.AuthClient) *KeeperClient {
	conn, err := grpc.Dial(
		cfg.GRPCServAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptors.NewAuthInterceptor(authClient)),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := &KeeperClient{
		KeeperClient: pb.NewKeeperClient(conn),
		authClient:   authClient,
		conn:         conn,
	}

	return client
}

func (c *KeeperClient) SignUp(ctx context.Context, in *pb.AuthRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	resp, err := c.KeeperClient.SignUp(ctx, in)
	if err != nil {
		return nil, err
	}

	c.authClient.SetToken(resp.AccessToken)

	return resp, nil
}

func (c *KeeperClient) SignIn(ctx context.Context, in *pb.AuthRequest, opts ...grpc.CallOption) (*pb.AuthResponse, error) {
	resp, err := c.KeeperClient.SignIn(ctx, in)
	if err != nil {
		return nil, err
	}

	c.authClient.SetToken(resp.AccessToken)

	return resp, nil
}

func (c *KeeperClient) Close() error {
	return c.conn.Close()
}
