package config

type Config struct {
	DBAddr       string
	GRPCServAddr string
	LogLevel     string
}

var defaultConfig = &Config{
	LogLevel:     "info",
	GRPCServAddr: ":3200",
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
		if c.DBAddr != "" {
			cfg.DBAddr = c.DBAddr
		}

		if c.GRPCServAddr != "" {
			cfg.GRPCServAddr = c.GRPCServAddr
		}

		if c.LogLevel != "" {
			cfg.LogLevel = c.LogLevel
		}
	}
}
