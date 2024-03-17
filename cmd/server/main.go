package main

import (
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/config"
	"github.com/Galish/goph-keeper/pkg/logger"
)

func main() {
	logger.Init()

	cfg := config.New()

	fmt.Println("Config:", cfg)

	logger.SetLevel(cfg.LogLevel)
}
