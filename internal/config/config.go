// Package config provides configuration management functionality for the gsync application.
// It handles loading, storing, and retrieving configuration settings from environment variables.
package config

import "os"

// Config holds application configuration settings.
type Config struct {
	// Role defines the application mode (e.g., "CLIENT" or "SERVER").
	Role string
}

var globalConfig *Config

// GetConfig returns the global configuration object.
// If the configuration hasn't been loaded yet, it loads it first.
func GetConfig() *Config {
	if globalConfig == nil {
		globalConfig = loadConfig()
	}
	return globalConfig
}

// loadConfig creates a new configuration instance and populates it
// with values from environment variables or defaults.
func loadConfig() *Config {
	config := &Config{}

	role := os.Getenv("ROLE")
	if role == "" {
		role = "CLIENT"
	}
	config.Role = role

	return config
}
