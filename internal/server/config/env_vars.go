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

	if authKey := os.Getenv("SECRET_KEY"); authKey != "" {
		c.AuthSecretKey = authKey
	}

	if certPath := os.Getenv("CERTIFICATE_PATH"); certPath != "" {
		c.CertPath = certPath
	}

	if keyPath := os.Getenv("PRIVATE_KEY_PATH"); keyPath != "" {
		c.KeyPath = keyPath
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		c.LogLevel = logLevel
	}
}
