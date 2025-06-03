package config

import "os"

type Config struct {
	DatabaseURL string
	Port        string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABSE_URL", "postgres://bhagavatam:password@localhost:5432/bhagavatam?sslMode=disable"),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
