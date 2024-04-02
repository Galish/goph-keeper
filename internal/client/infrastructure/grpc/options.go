package grpc

import (
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func withTransport(cfg *config.Config) grpc.DialOption {
	creds, err := credentials.NewClientTLSFromFile(cfg.CertPath, "")
	if err != nil {
		logger.WithError(err).Debug("error initializing client credentials")

		return grpc.WithTransportCredentials(insecure.NewCredentials())
	}

	return grpc.WithTransportCredentials(creds)
}
