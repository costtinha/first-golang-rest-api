package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppEnv  string
	AppPort string

	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	DBSSL    string
	LogLevel string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppEnv:  getEnv("APP_ENV", "dev"),
		AppPort: getEnv("APP_PORT", "5432"),
		DBHost:  getEnv("DB_HOST", "localhost"),
		DBPort:  getEnv("DB_Port", "5432"),
		DBUser:  getEnv("DB_USER", "postgres"),
		DBPass:  getEnv("DB_PASS", "postgres"),
		DBName:  getEnv("DB_NAME", "appdb"),
		DBSSL:   getEnv("DB_SSLMODE", "disable"),

		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
	return cfg, nil
}

func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s name=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.DBName, c.DBSSL)
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
