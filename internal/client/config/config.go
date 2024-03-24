package config

type Config struct {
	GRPCServAddr string
	LogLevel     string
}

func New() *Config {
	return &Config{
		GRPCServAddr: ":3200",
	}
}
