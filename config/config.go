package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

// menyimpan config API server
type ApiConfig struct {
	ApiPort string
	ApiHost string
}

// Menyimpan config Database
type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Driver   string
}

type Config struct {
	API ApiConfig
	DB  DbConfig
}

// Inisialisasi instance config
func NewConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error loading config .env file")
	}

	config := &Config{
		API: ApiConfig{
			ApiPort: os.Getenv("API_PORT"),
			ApiHost: os.Getenv("API_HOST"),
		},
		DB: DbConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Driver:   os.Getenv("DB_DRIVER"),
		},
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// validate apakah semua field config yang diperlukan telah diisi
func (c *Config) validate() error {
	if c.API.ApiPort == "" || c.DB.Driver == "" || c.DB.Host == "" || c.DB.Name == "" || c.DB.User == "" {
		return errors.New("All environment variables required")
	}
	return nil
}
