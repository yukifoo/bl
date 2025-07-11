package config

import (
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
	viper.AddConfigPath("$HOME/.config/bl")
	viper.AddConfigPath("/etc/bl")
	
	// Environment variables
	viper.SetEnvPrefix("BL")
	viper.AutomaticEnv()
	
	var config Config
	if err := viper.ReadInConfig(); err != nil {
		// Config file not found; create default
		return &config, nil
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
	
	configDir := filepath.Join("$HOME", ".config", "bl")
	configPath := filepath.Join(configDir, "config.yaml")
	
	return viper.WriteConfigAs(configPath)
}