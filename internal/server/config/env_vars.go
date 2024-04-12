package config

import "os"

func parseEnvVars(c *Config) {
	if dbAddr := os.Getenv("DATABASE_DSN"); dbAddr != "" {
		c.DBAddr = dbAddr
	}

	if grpcAddr := os.Getenv("GRPC_ADDRESS"); grpcAddr != "" {
		c.GRPCServAddr = grpcAddr
	}

	if authKey := os.Getenv("SECRET_KEY"); authKey != "" {
		c.AuthSecretKey = authKey
	}

	if passphrase := os.Getenv("PASSPHRASE"); passphrase != "" {
		c.EncryptPassphrase = passphrase
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
