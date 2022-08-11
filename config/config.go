package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string `json:"appport"`
}

func (c *Config) LoadConfig() error {
	// Implement viper load config
	c.AppPort = fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	return nil
}
