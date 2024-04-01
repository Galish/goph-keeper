package main

import (
	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
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

	appUI := ui.New()

	app := cli.NewApp(
		appUI,
		authManager,
		user.New(grpcClient),
		keeper.New(grpcClient),
	)

	sd := shutdowner.New(grpcClient, appUI)

	go app.Run()

	sd.Wait()
}
