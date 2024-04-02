package config

import "os"

func parseEnvVars(c *Config) {
	if grpcAddr := os.Getenv("GRPC_ADDRESS"); grpcAddr != "" {
		c.GRPCServAddr = grpcAddr
	}

	if certPath := os.Getenv("CERTIFICATE_PATH"); certPath != "" {
		c.CertPath = certPath
	}

	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		c.LogLevel = logLevel
	}
}