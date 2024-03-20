package config

import "flag"

func parseFlags(c *Config) {
	flag.StringVar(&c.DBAddr, "d", "", "DB address")
	flag.StringVar(&c.GRPCServAddr, "g", "", "GRPC server address")
	flag.StringVar(&c.LogLevel, "l", "", "Log level")
	flag.Parse()
}
