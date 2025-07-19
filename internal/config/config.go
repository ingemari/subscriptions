package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBHost     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		slog.Error("No .env file found")
	} else {
		slog.Info("Loaded .env file")
	}
	return &Config{
		ServerPort: os.Getenv("SERVER_PORT"),
		DBPort:     os.Getenv("DATABASE_PORT"),
		DBUser:     os.Getenv("DATABASE_USER"),
		DBName:     os.Getenv("DATABASE_NAME"),
		DBPassword: os.Getenv("DATABASE_PASSWORD"),
		DBHost:     os.Getenv("DATABASE_HOST"),
	}
}

func MakeDSN(cfg Config) string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBHost,
		cfg.DBPort,
	)
}
