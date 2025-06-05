package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Port      string
	BaseURL   string
	Token     string
	Timeout   time.Duration
	JWTSecret string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		Port:      getEnv("PORT", "8080"),
		BaseURL:   getEnv("DADATA_BASE_URL", "http://localhost:8080"),
		Token:     getEnv("DADATA_TOKEN", ""),
		Timeout:   getDurationEnv("TIMEOUT", 5*time.Second),
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	value := getEnv(key, "")
	if value == "" {
		return fallback
	}

	d, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Error parsing duration '%s': %s", value, err)
		return fallback
	}
	return d
}
