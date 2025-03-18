package internal

import (
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	// General settings
	AppName        string
	Version        string
	LogLevel       string
	ConfigFile     string
	LastUpdateCheck time.Time

	// User settings
	Username string
	Email    string

	// Path settings
	DataDir     string
	LogDir      string
	TemplateDir string
}

// NewConfig creates a new configuration instance with default values
func NewConfig() *Config {
	return &Config{
		AppName:  "{{.ProjectName}}",
		Version:  "0.1.0",
		LogLevel: "info",
	}
}

// LoadConfig loads the configuration from disk
func LoadConfig(configFile string) (*Config, error) {
	config := NewConfig()
	
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// Set default config locations
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		viper.SetConfigName(".{{.ProjectName}}")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(homeDir)
		viper.AddConfigPath(".")
	}

	// Set default values
	viper.SetDefault("AppName", config.AppName)
	viper.SetDefault("Version", config.Version)
	viper.SetDefault("LogLevel", config.LogLevel)
	
	// Try to read config file
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if we don't find a config file
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// Parse config into struct
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	// Set some computed defaults if not specified
	if config.DataDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		config.DataDir = filepath.Join(homeDir, ".{{.ProjectName}}", "data")
	}

	if config.LogDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		config.LogDir = filepath.Join(homeDir, ".{{.ProjectName}}", "logs")
	}

	return config, nil
}

// SaveConfig saves the current configuration to disk
func (c *Config) SaveConfig() error {
	viper.Set("AppName", c.AppName)
	viper.Set("Version", c.Version)
	viper.Set("LogLevel", c.LogLevel)
	viper.Set("LastUpdateCheck", c.LastUpdateCheck)
	viper.Set("Username", c.Username)
	viper.Set("Email", c.Email)
	viper.Set("DataDir", c.DataDir)
	viper.Set("LogDir", c.LogDir)
	viper.Set("TemplateDir", c.TemplateDir)

	return viper.WriteConfig()
} 