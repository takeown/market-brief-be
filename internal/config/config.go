package config

import (
	"os"
)

type Config struct {
	Env  string
	Port string
}

func Load() *Config {
	return &Config{
		Env:  getEnv("APP_ENV", "development"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
