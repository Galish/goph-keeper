package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Galish/goph-keeper/internal/client/auth"
	"github.com/Galish/goph-keeper/internal/client/cli"
	"github.com/Galish/goph-keeper/internal/client/config"
	"github.com/Galish/goph-keeper/internal/client/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/client/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/client/usecase/user"
)

func main() {
	cfg := config.New()

	authClient := auth.New()

	grpcClient := grpc.NewClient(cfg, authClient)
	defer grpcClient.Close()

	userUsecase := user.New(grpcClient)
	keeperUsecase := keeper.New(grpcClient)

	app := cli.NewApp(authClient, userUsecase, keeperUsecase)
	go app.Run()

	sigs := make(chan os.Signal, 1)

	signal.Notify(
		sigs,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	// go func() {
	<-sigs
	app.Close()
	// }()
}
