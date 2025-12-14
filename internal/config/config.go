package config

// ServerConfig:포트, 환경 설정 
// DatabaseConfig: DB드라이버,DSN(SQlite 파일 경로)
// AuthConfig: 세션 시크릿, 세션 유효기간
// Load(): 환경변수 읽어서 Config 구조체 생성

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Auth     AuthConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Env  string // development, production
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Driver string // sqlite, postgres (future)
	DSN    string // data source name (file path for sqlite)
}

// AuthConfig holds authentication configuration
type AuthConfig struct {
	SessionSecret string
	SessionMaxAge int // in seconds
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Env:  getEnv("SERVER_ENV", "development"),
		},
		Database: DatabaseConfig{
			Driver: getEnv("DB_DRIVER", "sqlite"),
			DSN:    getEnv("DB_DSN", "./data/blog.db"),
		},
		Auth: AuthConfig{
			SessionSecret: getEnv("SESSION_SECRET", "change-me-in-production"),
			SessionMaxAge: getEnvAsInt("SESSION_MAX_AGE", 86400), // 24 hours
		},
	}

	// Validate required fields
	if cfg.Auth.SessionSecret == "change-me-in-production" && cfg.Server.Env == "production" {
		return nil, fmt.Errorf("SESSION_SECRET must be set in production")
	}

	return cfg, nil
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt gets an environment variable as int or returns a default value
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	
	return value
}
