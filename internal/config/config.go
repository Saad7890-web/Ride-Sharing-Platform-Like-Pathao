package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppEnv  string
	HTTPPort string

	DB DBConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		AppName: os.Getenv("APP_NAME"),
		AppEnv:  os.Getenv("APP_ENV"),
		HTTPPort: os.Getenv("HTTP_PORT"),
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
			MaxOpenConns:    mustInt("DB_MAX_OPEN_CONNS"),
			MaxIdleConns:    mustInt("DB_MAX_IDLE_CONNS"),
			ConnMaxLifetime: mustDuration("DB_CONN_MAX_LIFETIME"),
		},
	}

	return cfg
}

func mustInt(key string) int {
	val := os.Getenv(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("invalid int for %s", key)
	}
	return i
}

func mustDuration(key string) time.Duration {
	val := os.Getenv(key)
	d, err := time.ParseDuration(val)
	if err != nil {
		log.Fatalf("invalid duration for %s", key)
	}
	return d
}
