package grpc

import (
	"log"
	"net"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// KeeperServer represents gRPC server.
type KeeperServer struct {
	pb.UnimplementedKeeperServer

	cfg    *config.Config
	user   usecase.User
	keeper usecase.Keeper
	server *grpc.Server
}

// NewServer configures and creates a gRPC server.
func NewServer(
	cfg *config.Config,
	user usecase.User,
	keeper usecase.Keeper,
) *KeeperServer {
	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptors.NewAuthInterceptor(user).Unary(),
			interceptors.LoggerInterceptor,
		),
	)
	reflection.Register(s)

	server := &KeeperServer{
		cfg:    cfg,
		server: s,
		user:   user,
		keeper: keeper,
	}

	pb.RegisterKeeperServer(s, server)

	return server
}

// Run listens and serves GRPC requests.
func (s *KeeperServer) Run() error {
	listener, err := net.Listen("tcp", s.cfg.GRPCServAddr)
	if err != nil {
		log.Fatal(err)
	}

	logger.Info("running gRPC server on ", s.cfg.GRPCServAddr)

	return s.server.Serve(listener)
}

// Close is executed to release the memory
func (s *KeeperServer) Close() error {
	logger.Info("shutting down the gRPC server")

	s.server.GracefulStop()

	return nil
}
