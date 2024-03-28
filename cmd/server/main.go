package main

import (
	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/server/repository/psql"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/auth"
	"github.com/Galish/goph-keeper/pkg/logger"
)

func main() {
	logger.Init()

	cfg := config.New()

	logger.SetLevel(cfg.LogLevel)

	repo, err := psql.New(cfg)
	if err != nil {
		panic(err)
	}

	jwtManager := auth.NewJWTManager("secret_key")

	userUsecase := user.New(repo, jwtManager)
	keeperUsecase := keeper.New(repo)

	grpcServer := grpc.NewServer(cfg, userUsecase, keeperUsecase)
	if err := grpcServer.Run(); err != nil {
		panic(err)
	}
}
