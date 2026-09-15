package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("Database is required")
	}

	return cfg, nil
}
