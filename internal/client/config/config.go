package config

type Config struct {
	GRPCServAddr string
	CertPath     string
	LogLevel     string
}

var defaultConfig = &Config{
	GRPCServAddr: ":3200",
	LogLevel:     "info",
}

func New() *Config {
	var flags = new(Config)
	var envVars = new(Config)

	parseFlags(flags)
	parseEnvVars(envVars)

	return initConfig(
		withConfig(defaultConfig),
		withConfig(flags),
		withConfig(envVars),
	)
}

func initConfig(opts ...func(*Config)) *Config {
	cfg := &Config{}

	for _, o := range opts {
		o(cfg)
	}

	return cfg
}

func withConfig(c *Config) func(*Config) {
	return func(cfg *Config) {
		if c.GRPCServAddr != "" {
			cfg.GRPCServAddr = c.GRPCServAddr
		}

		if c.CertPath != "" {
			cfg.CertPath = c.CertPath
		}

		if c.LogLevel != "" {
			cfg.LogLevel = c.LogLevel
		}
	}
}
