package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB        DBConfig
	JWT       JWTConfig
	Server    ServerConfig
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret          string
	Expiry          int // hours
	DefaultPassword string
}

type ServerConfig struct {
	Port string
	Mode string
}

var AppConfig Config

func LoadConfig() {
	// Load .env file (ignore error if not found)
	_ = godotenv.Load()

	AppConfig = Config{
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "3306"),
			User:     getEnv("DB_USER", "root"),
			Password: getEnv("DB_PASSWORD", ""),
			Name:     getEnv("DB_NAME", "isdc_db"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "default-secret-change-me"),
			Expiry:          getEnvAsInt("JWT_EXPIRY", 24),
			DefaultPassword: getEnv("DEFAULT_PASSWORD", "password321!*"),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "debug"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
