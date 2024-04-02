package grpc

import (
	"log"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc"
)

type KeeperClient struct {
	pb.KeeperClient
	auth *auth.AuthManager
	conn *grpc.ClientConn
}

func NewClient(cfg *config.Config, auth *auth.AuthManager) *KeeperClient {
	conn, err := grpc.Dial(
		cfg.GRPCServAddr,
		grpc.WithUnaryInterceptor(interceptors.NewAuthInterceptor(auth)),
		withTransport(cfg),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := &KeeperClient{
		KeeperClient: pb.NewKeeperClient(conn),
		auth:         auth,
		conn:         conn,
	}

	return client
}

func (c *KeeperClient) Close() error {
	logger.Info("shutting down the gRPC client")

	return c.conn.Close()
}
