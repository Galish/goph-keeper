package main

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/server/repository/psql"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/logger"
)

func main() {
	logger.Init()

	cfg := config.New()

	fmt.Println("Config:", cfg)

	logger.SetLevel(cfg.LogLevel)

	repo, err := psql.New(cfg)
	if err != nil {
		panic(err)
	}

	userUsecase := user.New(repo, "secret_key")
	keeperUsecase := keeper.New(repo)

	grpcServer := grpc.NewServer(cfg, userUsecase, keeperUsecase)
	if err := grpcServer.Run(); err != nil {
		panic(err)
	}
}
