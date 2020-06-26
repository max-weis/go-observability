package config

import (
	"go.uber.org/zap"
	"os"
)

type Config struct {
	Port   string
	logger zap.Logger
}

func NewConfig(logger zap.Logger) Config {
	var c Config

	c.logger = logger
	c.getEnvVar()

	return c
}

func (c *Config) getEnvVar() {
	port := os.Getenv("PORT")
	if port == "" {
		c.logger.Fatal("Could not read env var")
	}

	c.Port = port
}
