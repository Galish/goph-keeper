package config

import "flag"

func parseFlags(c *Config) {
	flag.StringVar(&c.DBAddr, "d", "", "DB address")
	flag.StringVar(&c.GRPCServAddr, "g", "", "gRPC server address")
	flag.StringVar(&c.AuthSecretKey, "s", "", "string used to sign the JWT token")
	flag.StringVar(&c.EncryptPassphrase, "p", "", "encryption passphrase")
	flag.StringVar(&c.CertPath, "c", "", "certificate file path")
	flag.StringVar(&c.KeyPath, "k", "", "private key file path")
	flag.StringVar(&c.LogLevel, "l", "", "Log level")
	flag.Parse()
}
