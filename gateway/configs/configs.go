package configs

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	
	Server struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Endpoints struct {
		Auth string 
	}

	Environment string
}

var cfg *Config

func GetConfig() *Config {
	return cfg	
}

func Load() {
	godotenv.Load("./.env")

	cfg = &Config{}

	cfg.Server.Port = getEnv("GATEWAY_SERVER_PORT", "8000")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 15

	cfg.Endpoints.Auth = getEnv("AUTH_ENDPOINT", "")

	cfg.Environment = getEnv("ENV", "development")	
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
