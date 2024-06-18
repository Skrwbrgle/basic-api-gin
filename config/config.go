package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiPort    string
	ApiHost    string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbDriver   string
}

// Inisialisasi instance config
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, errors.New("error loading config .env file")
	}

	config := &Config{
		ApiPort:    os.Getenv("API_PORT"),
		ApiHost:    os.Getenv("API_HOST"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		DbDriver:   os.Getenv("DB_DRIVER"),
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// validate apakah semua field config yang diperlukan telah diisi
func (c *Config) validate() error {
	if c.ApiPort == "" || c.DbDriver == "" || c.DbHost == "" || c.DbName == "" || c.DbUser == "" {
		return errors.New("all environment variables required")
	}
	return nil
}

func (c *Config) DbConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName)
}
