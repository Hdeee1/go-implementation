package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost		string
	DBPort		int
	DBUser		string
	DBPassword	string
	DBName		string
	ServerPort	string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	dbPort := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(dbPort)
	if err != nil {
		return nil, err
	}

	config :=  Config{
		DBHost: os.Getenv("DB_HOST"),
		DBPort: port,
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}

	return &config, nil
}

func (cfg *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
}