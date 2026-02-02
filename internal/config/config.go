package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env            string
	Port           string
	FinnhubAPIKey  string
	FinnhubBaseURL string
}

func Load() *Config {
	// .env 파일 로드 (파일이 없어도 에러 무시)
	_ = godotenv.Load()

	return &Config{
		Env:            getEnv("APP_ENV", "development"),
		Port:           getEnv("PORT", "8080"),
		FinnhubAPIKey:  getEnv("FINNHUB_API_KEY", ""),
		FinnhubBaseURL: getEnv("FINNHUB_BASE_URL", "https://finnhub.io/api/v1"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
