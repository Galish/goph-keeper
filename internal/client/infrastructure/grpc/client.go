// Package implements gRPC client.
package grpc

import (
	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc"
)

// KeeperClient represents the gRPC client.
type KeeperClient struct {
	pb.KeeperClient
	auth *auth.Manager
	conn *grpc.ClientConn
}

// NewClient returns a new gRPC client.
func NewClient(cfg *config.Config, auth *auth.Manager) *KeeperClient {
	conn, err := grpc.Dial(
		cfg.GRPCServAddr,
		grpc.WithUnaryInterceptor(interceptors.NewAuthInterceptor(auth)),
		withTransport(cfg),
	)
	if err != nil {
		logger.Fatal(err)
	}

	client := &KeeperClient{
		KeeperClient: pb.NewKeeperClient(conn),
		auth:         auth,
		conn:         conn,
	}

	return client
}

// Close closes the gRPC connection.
func (c *KeeperClient) Close() error {
	logger.Info("shutting down the gRPC client")

	return c.conn.Close()
}
