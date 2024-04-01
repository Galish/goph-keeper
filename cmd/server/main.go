package main

import (
	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/internal/server/infrastructure/grpc"
	"github.com/Galish/goph-keeper/internal/server/repository/psql"
	"github.com/Galish/goph-keeper/internal/server/usecase/keeper"
	"github.com/Galish/goph-keeper/internal/server/usecase/user"
	"github.com/Galish/goph-keeper/pkg/auth"
	"github.com/Galish/goph-keeper/pkg/logger"
	"github.com/Galish/goph-keeper/pkg/shutdowner"
)

func main() {
	logger.Init()

	cfg := config.New()

	logger.SetLevel(cfg.LogLevel)

	store, err := psql.New(cfg)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer(
		cfg,
		user.New(store, auth.NewJWTManager(cfg.AuthSecretKey)),
		keeper.New(store),
	)

	sd := shutdowner.New(grpcServer, store)

	go func() {
		if err := grpcServer.Run(); err != nil {
			panic(err)
		}
	}()

	sd.Wait()
}
