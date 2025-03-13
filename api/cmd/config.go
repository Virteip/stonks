package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Service struct {
		Name string `json:"name"`
	} `json:"service"`

	Database struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Name     string `json:"name"`
		Options  string `json:"options"`
	} `json:"database"`

	Server struct {
		Host          string `json:"host"`
		Port          int    `json:"port"`
		APIKey        string `json:"APIKey"`
		AllowedOrigin string `json:"allowedOrigin"`
	} `json:"server"`

	ExternalStocksAPI struct {
		URL        string `json:"url"`
		AuthHeader string `json:"authHeader"`
		AuthToken  string `json:"authToken"`
	} `json:"externalStocksAPI"`
}

func LoadConfig(environment string) (*Config, error) {
	if environment == "local" {
		return loadFromFile("configs/local.config.json")
	}

	// For all other environments, load from environment variables
	return loadFromEnv()
}

// Load configuration from file
func loadFromFile(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file %s: %v", configPath, err)
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}

	return config, nil
}

// Load configuration from environment variables
func loadFromEnv() (*Config, error) {
	config := &Config{}

	// Service config
	serviceName, err := getRequiredEnv("SERVICE_NAME")
	if err != nil {
		return nil, err
	}
	config.Service.Name = serviceName

	// Database config
	dbUser, err := getRequiredEnv("DB_USER")
	if err != nil {
		return nil, err
	}
	config.Database.User = dbUser

	dbPassword, err := getRequiredEnv("DB_PASSWORD")
	if err != nil {
		return nil, err
	}
	config.Database.Password = dbPassword

	dbHost, err := getRequiredEnv("DB_HOST")
	if err != nil {
		return nil, err
	}
	config.Database.Host = dbHost

	dbName, err := getRequiredEnv("DB_NAME")
	if err != nil {
		return nil, err
	}
	config.Database.Name = dbName

	dbOptions, err := getRequiredEnv("DB_OPTIONS")
	if err != nil {
		return nil, err
	}
	config.Database.Options = dbOptions

	dbPortStr, err := getRequiredEnv("DB_PORT")
	if err != nil {
		return nil, err
	}
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %v", err)
	}
	config.Database.Port = dbPort

	// Server config
	serverAllowedOrigin, err := getRequiredEnv("SERVER_ALLOWED_ORIGIN")
	if err != nil {
		return nil, err
	}
	config.Server.AllowedOrigin = serverAllowedOrigin

	serverHost, err := getRequiredEnv("SERVER_HOST")
	if err != nil {
		return nil, err
	}
	config.Server.Host = serverHost

	serverAPIKey, err := getRequiredEnv("SERVER_API_KEY")
	if err != nil {
		return nil, err
	}
	config.Server.APIKey = serverAPIKey

	serverPortStr, err := getRequiredEnv("SERVER_PORT")
	if err != nil {
		return nil, err
	}
	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %v", err)
	}
	config.Server.Port = serverPort

	// External API config
	apiURL, err := getRequiredEnv("API_URL")
	if err != nil {
		return nil, err
	}
	config.ExternalStocksAPI.URL = apiURL

	apiAuthHeader, err := getRequiredEnv("API_AUTH_HEADER")
	if err != nil {
		return nil, err
	}
	config.ExternalStocksAPI.AuthHeader = apiAuthHeader

	apiAuthToken, err := getRequiredEnv("API_AUTH_TOKEN")
	if err != nil {
		return nil, err
	}
	config.ExternalStocksAPI.AuthToken = apiAuthToken

	return config, nil
}

// Helper function to get a required environment variable
func getRequiredEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("required environment variable %s is not set", key)
	}
	return value, nil
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.Options)
}

func (c *Config) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
