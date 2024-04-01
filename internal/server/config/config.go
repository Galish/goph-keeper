package config

type Config struct {
	DBAddr        string
	DBInitPath    string
	GRPCServAddr  string
	AuthSecretKey string
	LogLevel      string
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
		if c.DBAddr != "" {
			cfg.DBAddr = c.DBAddr
		}

		if c.DBInitPath != "" {
			cfg.DBInitPath = c.DBInitPath
		}

		if c.GRPCServAddr != "" {
			cfg.GRPCServAddr = c.GRPCServAddr
		}

		if c.AuthSecretKey != "" {
			cfg.AuthSecretKey = c.AuthSecretKey
		}

		if c.LogLevel != "" {
			cfg.LogLevel = c.LogLevel
		}
	}
}
