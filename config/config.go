// Package config provides configuration management for gascity.
// It supports loading configuration from environment variables and config files.
package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration values for the application.
type Config struct {
	// Server configuration
	Host string
	Port int

	// Database configuration
	DatabaseURL string

	// Application configuration
	LogLevel string
	Environment string

	// Gas price configuration
	GasAPIEndpoint string
	GasAPIKey      string
	PollIntervalSec int
}

// Load reads configuration from environment variables and returns a Config.
// Required environment variables will cause an error if not set.
func Load() (*Config, error) {
	cfg := &Config{
		Host:            getEnv("HOST", "0.0.0.0"),
		Port:            getEnvAsInt("PORT", 8080),
		DatabaseURL:     getEnv("DATABASE_URL", ""),
		LogLevel:        getEnv("LOG_LEVEL", "info"),
		Environment:     getEnv("ENVIRONMENT", "development"),
		GasAPIEndpoint:  getEnv("GAS_API_ENDPOINT", "https://api.etherscan.io/api"),
		GasAPIKey:       getEnv("GAS_API_KEY", ""),
		PollIntervalSec: getEnvAsInt("POLL_INTERVAL_SEC", 15),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// validate checks that all required configuration values are present.
func (c *Config) validate() error {
	if c.GasAPIKey == "" {
		return fmt.Errorf("GAS_API_KEY is required")
	}
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("PORT must be between 1 and 65535, got %d", c.Port)
	}
	if c.PollIntervalSec < 1 {
		return fmt.Errorf("POLL_INTERVAL_SEC must be at least 1, got %d", c.PollIntervalSec)
	}
	return nil
}

// Addr returns the full address string for the HTTP server.
func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// IsDevelopment returns true if the environment is development.
func (c *Config) IsDevelopment() bool {
	return c.Environment == "development"
}

// getEnv retrieves an environment variable or returns a default value.
func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

// getEnvAsInt retrieves an environment variable as an integer or returns a default value.
func getEnvAsInt(key string, defaultVal int) int {
	strVal := getEnv(key, "")
	if strVal == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(strVal)
	if err != nil {
		return defaultVal
	}
	return val
}
