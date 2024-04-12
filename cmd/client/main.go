package main

import (
	"context"

	"github.com/Galish/goph-keeper/internal/client/app"
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	healthcheck "github.com/Galish/goph-keeper/internal/client/usecase/health_check"
	"github.com/Galish/goph-keeper/internal/client/usecase/notes"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
	"github.com/Galish/goph-keeper/pkg/logger"
	"github.com/Galish/goph-keeper/pkg/shutdowner"
)

func main() {
	logger.Init()

	cfg := config.New()

	logger.SetLevel(cfg.LogLevel)

	authManager := auth.New()

	grpcClient := grpc.NewClient(cfg, authManager)

	app := app.New(
		authManager,
		user.New(grpcClient),
		notes.New(grpcClient),
		healthcheck.New(grpcClient),
	)

	sd := shutdowner.New(grpcClient, app)

	go app.Run(context.Background())

	sd.Wait()
}
