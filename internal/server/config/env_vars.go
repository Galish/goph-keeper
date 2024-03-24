package config

import "os"

func parseEnvVars(c *Config) {
	if dbAddr := os.Getenv("DATABASE_DSN"); dbAddr != "" {
		c.DBAddr = dbAddr
	}

	if dbInitPath := os.Getenv("DATABASE_INIT_PATH"); dbInitPath != "" {
		c.DBInitPath = dbInitPath
	}

	if grpcAddr := os.Getenv("GRPC_ADDRESS"); grpcAddr != "" {
		c.GRPCServAddr = grpcAddr
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		c.LogLevel = logLevel
	}
}
