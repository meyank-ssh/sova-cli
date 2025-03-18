package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config contains all configuration for the application
type Config struct {
	// Server configuration
	ServerPort int
	ServerHost string
	
	// Database configuration
	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbSSLMode  string
	
	// JWT configuration
	JWTSecret     string
	JWTExpiration time.Duration
	
	// Environment
	Environment string
	
	// Application
	AppName    string
	AppVersion string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() (*Config, error) {
	// Load .env file if it exists
	godotenv.Load()
	
	// Set defaults
	config := &Config{
		ServerPort:    8080,
		ServerHost:    "0.0.0.0",
		DbHost:        "localhost",
		DbPort:        5432,
		DbUser:        "postgres",
		DbPassword:    "postgres",
		DbName:        "{{.ProjectName}}",
		DbSSLMode:     "disable",
		JWTSecret:     "your-secret-key",
		JWTExpiration: 24 * time.Hour,
		Environment:   "development",
		AppName:       "{{.ProjectName}}",
		AppVersion:    "0.1.0",
	}
	
	// Override with environment variables if set
	if port, err := strconv.Atoi(getEnv("SERVER_PORT", "8080")); err == nil {
		config.ServerPort = port
	}
	
	config.ServerHost = getEnv("SERVER_HOST", config.ServerHost)
	
	config.DbHost = getEnv("DB_HOST", config.DbHost)
	if port, err := strconv.Atoi(getEnv("DB_PORT", "5432")); err == nil {
		config.DbPort = port
	}
	config.DbUser = getEnv("DB_USER", config.DbUser)
	config.DbPassword = getEnv("DB_PASSWORD", config.DbPassword)
	config.DbName = getEnv("DB_NAME", config.DbName)
	config.DbSSLMode = getEnv("DB_SSL_MODE", config.DbSSLMode)
	
	config.JWTSecret = getEnv("JWT_SECRET", config.JWTSecret)
	if exp, err := strconv.Atoi(getEnv("JWT_EXPIRATION_HOURS", "24")); err == nil {
		config.JWTExpiration = time.Duration(exp) * time.Hour
	}
	
	config.Environment = getEnv("ENVIRONMENT", config.Environment)
	
	return config, nil
}

// GetDatabaseURL returns the database connection URL
func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName, c.DbSSLMode)
}

// IsProduction returns true if the environment is production
func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}

// IsDevelopment returns true if the environment is development
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development" || c.Environment == ""
}

// getEnv gets an environment variable or returns the default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
} 