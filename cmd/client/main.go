package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli"
	"github.com/Galish/goph-keeper/internal/client/cli/ui"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
)

func main() {
	cfg := config.New()

	authManager := auth.New()

	grpcClient := grpc.NewClient(cfg, authManager)

	appUI := ui.New()

	app := cli.NewApp(
		appUI,
		authManager,
		user.New(grpcClient),
		keeper.New(grpcClient),
	)
	go app.Run()

	sigs := make(chan os.Signal, 1)

	signal.Notify(
		sigs,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	<-sigs
	grpcClient.Close()
	app.Close()
	appUI.Close()
}
