package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Address string
	Driver  string
	DB      string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load("../../config/.env"); err != nil {
		return nil, errors.New("failed to load environment variables")
	}

	config := &Config{
		Address: os.Getenv("ADDRESS"),
		Driver: os.Getenv("SQLDRIVER"),
		DB: os.Getenv("SQLSOURCE"),
	}

	return config, nil
}