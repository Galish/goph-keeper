package grpc

import (
	"crypto/tls"
	"log"
	"net"

	pb "github.com/Galish/goph-keeper/api/proto"
	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc/interceptors"
	"github.com/Galish/goph-keeper/internal/server/usecase"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

// KeeperServer represents gRPC server.
type KeeperServer struct {
	pb.UnimplementedKeeperServer

	cfg    *config.Config
	user   usecase.User
	notes  usecase.SecureNotes
	server *grpc.Server
}

// NewServer configures and creates a gRPC server.
func NewServer(cfg *config.Config, user usecase.User, notes usecase.SecureNotes) *KeeperServer {
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptors.NewAuthInterceptor(user).Unary(),
			interceptors.LoggerInterceptor,
		),
	}

	cert, err := tls.LoadX509KeyPair(cfg.CertPath, cfg.KeyPath)
	if err != nil {
		logger.WithError(err).Debug("error initializing server credentials")
	} else {
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}

	s := grpc.NewServer(opts...)
	reflection.Register(s)

	server := &KeeperServer{
		cfg:    cfg,
		server: s,
		user:   user,
		notes:  notes,
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

// Close is executed to release the memory.
func (s *KeeperServer) Close() error {
	logger.Info("shutting down the gRPC server")

	s.server.GracefulStop()

	return nil
}
