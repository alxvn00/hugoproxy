package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	DBConfig
	ServerConfig
	AuthConfig
}

type DBConfig struct {
	Host     string
	PortDB   string
	Username string
	Password string
	Name     string
}

type ServerConfig struct {
	Port    string
	BaseURL string
	Timeout time.Duration
}

type AuthConfig struct {
	JWTSecret string
	Token     string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		// without logs
	}

	log.Println("=== ACTUAL BUILD ===")

	return &Config{
		DBConfig: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			PortDB:   getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "geo_users"),
		},
		ServerConfig: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			BaseURL: getEnv("DADATA_BASE_URL", "http://localhost:8080"),
			Timeout: getDurationEnv("TIMEOUT", 5*time.Second),
		},
		AuthConfig: AuthConfig{
			JWTSecret: getEnv("JWT_SECRET", ""),
			Token:     getEnv("DADATA_TOKEN", ""),
		},
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
