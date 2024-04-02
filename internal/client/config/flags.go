package config

import "flag"

func parseFlags(c *Config) {
	flag.StringVar(&c.GRPCServAddr, "g", "", "gRPC server address")
	flag.StringVar(&c.CertPath, "c", "", "certificate file path")
	flag.StringVar(&c.LogLevel, "l", "", "Log level")
	flag.Parse()
}
