// Package config handles configuration management for the Backlog CLI
package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config represents the application configuration
type Config struct {
	BaseURL string `mapstructure:"base_url"`
	APIKey  string `mapstructure:"api_key"`
	Space   string `mapstructure:"space"`
}

// Load loads configuration from file and environment
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Add config paths
	viper.AddConfigPath(".")

	// Get user home directory and add config path
	homeDir, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(filepath.Join(homeDir, ".config", "bl"))
	}
	viper.AddConfigPath("/etc/bl")

	// Environment variables
	viper.SetEnvPrefix("BL")
	viper.AutomaticEnv()

	var config Config
	if err := viper.ReadInConfig(); err != nil {
		// Check if it's a file not found error
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found; create default
			return &config, nil
		}
		// Other error occurred (permission, syntax, etc.)
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Save saves configuration to file
func (c *Config) Save() error {
	viper.Set("base_url", c.BaseURL)
	viper.Set("api_key", c.APIKey)
	viper.Set("space", c.Space)

	// Get user home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir := filepath.Join(homeDir, ".config", "bl")
	configPath := filepath.Join(configDir, "config.yaml")

	// Create config directory if it doesn't exist
	if err := os.MkdirAll(configDir, 0o750); err != nil {
		return err
	}

	return viper.WriteConfigAs(configPath)
}
