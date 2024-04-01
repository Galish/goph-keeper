package config

import "flag"

func parseFlags(c *Config) {
	flag.StringVar(&c.DBAddr, "d", "", "DB address")
	flag.StringVar(&c.DBInitPath, "i", "", "DB init file path")
	flag.StringVar(&c.GRPCServAddr, "g", "", "gRPC server address")
	flag.StringVar(&c.AuthSecretKey, "s", "", "string used to sign the JWT token")
	flag.StringVar(&c.LogLevel, "l", "", "Log level")
	flag.Parse()
}
