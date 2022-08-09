package config

type Config struct {
	AppPort string `json:"appport"`
}

func (c *Config) LoadConfig() error {
	// Implement viper load config
	c.AppPort = ":8080"
	return nil
}
