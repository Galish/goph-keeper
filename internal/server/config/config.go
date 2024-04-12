package config

import (
	"errors"

	"github.com/Galish/goph-keeper/pkg/logger"
)

type Config struct {
	DBAddr            string
	GRPCServAddr      string
	AuthSecretKey     string
	EncryptPassphrase string
	CertPath          string
	KeyPath           string
	LogLevel          string
}

var defaultConfig = &Config{
	GRPCServAddr:      ":3200",
	AuthSecretKey:     "secret_key",
	EncryptPassphrase: "pqssjyEpfbwxyAqTPJdP28ueaVmrjEjV",
	LogLevel:          "info",
}

func New() *Config {
	var (
		flags   = new(Config)
		envVars = new(Config)
	)

	parseFlags(flags)
	parseEnvVars(envVars)

	return initConfig(
		withConfig(defaultConfig),
		withConfig(flags),
		withConfig(envVars),
	)
}

func (c *Config) Validate() error {
	switch len(c.EncryptPassphrase) {
	case 16:
	case 24:
	case 32:
		break

	default:
		return errors.New("AES only supports key sizes of 16, 24 or 32 bytes")
	}

	return nil
}

func initConfig(opts ...func(*Config)) *Config {
	cfg := &Config{}

	for _, o := range opts {
		o(cfg)
	}

	if err := cfg.Validate(); err != nil {
		logger.Fatal(err)
	}

	return cfg
}

func withConfig(c *Config) func(*Config) {
	return func(cfg *Config) {
		if c.DBAddr != "" {
			cfg.DBAddr = c.DBAddr
		}

		if c.GRPCServAddr != "" {
			cfg.GRPCServAddr = c.GRPCServAddr
		}

		if c.AuthSecretKey != "" {
			cfg.AuthSecretKey = c.AuthSecretKey
		}

		if c.EncryptPassphrase != "" {
			cfg.EncryptPassphrase = c.EncryptPassphrase
		}

		if c.CertPath != "" {
			cfg.CertPath = c.CertPath
		}

		if c.KeyPath != "" {
			cfg.KeyPath = c.KeyPath
		}

		if c.LogLevel != "" {
			cfg.LogLevel = c.LogLevel
		}
	}
}
