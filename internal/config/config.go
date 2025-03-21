package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration settings.
type Config struct {
	Server struct {
		Port  string `env:"SERVER_PORT,default=8080"` // Example: Port from env or default 8080
		Debug bool   `env:"SERVER_DEBUG,default=false"`
	}
	Database struct {
		Host         string `env:"DATABASE_HOST,default=localhost"`
		Port         int    `env:"DATABASE_PORT,default=5432"` // Example: Integer port from env or default 5432
		Username     string `env:"DATABASE_USERNAME,default=yourusername"`
		Password     string `env:"DATABASE_PASSWORD,default=yourpassword"`
		DatabaseName string `env:"DATABASE_NAME,default=yourdatabase"`
	}
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Convert integer values from strings to integers
	portStr := os.Getenv("DATABASE_PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid DATABASE_PORT: %w", err)
		}
		cfg.Database.Port = port
	}

	// You can add more conversions here for other integer or boolean values.

	return cfg, nil
}
