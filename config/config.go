package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config represents the application configuration.
type Config struct {
	Port        string
	DatabaseURL string
}

// Load loads the configuration from environment variables.
func Load() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable is not set")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
