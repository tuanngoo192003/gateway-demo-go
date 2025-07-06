package config

import (
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Port         string
		Host         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}

	JWT struct {
		Secret        string
		Tokenexpiry   time.Duration
		RefreshExpiry time.Duration
	}

	S3 struct {
		Region    string
		Bucket    string
		AccessKey string
		SecretKey string
	}

	Environment string
}

var cfg *Config

func GetConfig() *Config {
	return cfg
}

func Load() (*Config, error) {
	godotenv.Load("./.env")

	cfg = &Config{}

	//Server config
	cfg.Server.Host = getEnv("AUTH_SERVER_HOST", "")
	cfg.Server.Port = getEnv("AUTH_SERVER_PORT", "8000")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 15

	//Database config
	cfg.Database.Host = getEnv("AUTH_DB_HOST", "AUTHdb")
	cfg.Database.Port = getEnv("AUTH_DB_PORT", "5432")
	cfg.Database.User = getEnv("AUTH_DB_USER", "postgres")
	cfg.Database.Password = getEnv("AUTH_DB_PASSWORD", "postgres")
	cfg.Database.DBName = getEnv("AUTH_DB_NAME", "userservicedb")
	cfg.Database.SSLMode = getEnv("AUTH_DB_SSLMODE", "disable")

	//JWT config
	cfg.JWT.Secret = getEnv("JWT_SECRET", "a-string-secret-at-least-256-bits-long")
	cfg.JWT.Tokenexpiry = time.Hour * 24
	cfg.JWT.RefreshExpiry = time.Hour * 168

	//S3 config
	cfg.S3.Region = getEnv("S3_REGION", "ap-southeast-2")
	cfg.S3.Bucket = getEnv("S3_BUCKET", "my-cinema-app-bucket")
	cfg.S3.AccessKey = getEnv("AWS_ACCESS_KEY", "")
	cfg.S3.SecretKey = getEnv("AWS_SECRET_KEY", "")

	cfg.Environment = getEnv("ENV", "development")

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.DBName,
		c.Database.SSLMode,
	)
}
