package grpc

import (
	"log"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/client/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type KeeperClient struct {
	cfg     *config.Config
	conn    *grpc.ClientConn
	service pb.KeeperClient
}

func NewClient(cfg *config.Config) *KeeperClient {
	conn, err := grpc.Dial(
		cfg.GRPCServAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := &KeeperClient{
		cfg:     cfg,
		conn:    conn,
		service: pb.NewKeeperClient(conn),
	}

	return client
}

func (c *KeeperClient) Close() error {
	return c.conn.Close()
}
