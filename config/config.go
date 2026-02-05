package config

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	Logger, _ = zap.NewProduction()
}

type Config struct {
	DBDSN     string
	JWTSecret string
	RedisAddr string
}

func Load() *Config {
	return &Config{
		DBDSN:     os.Getenv("DB_DSN"), // e.g., "user:pass@tcp(localhost:3306)/db"
		JWTSecret: os.Getenv("JWT_SECRET"),
		RedisAddr: os.Getenv("REDIS_ADDR"), // e.g., "localhost:6379"
	}
}
