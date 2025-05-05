package configuration

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	DOMAIN      string `yaml:"DOMAIN"`
	SERVER_PORT string `yaml:"SERVER_PORT"`
}

type Database struct {
	DB_HOST     string
	DB_PORT     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
}

type Configuration struct {
	SERVER   Server   `yaml:"SERVER"`
	DATABASE Database `yaml:"DATABASE"`
}

func Load(file string) (*Configuration, error) {
	err := godotenv.Load(file)
	if err != nil {
		fmt.Println(".env file not found. Operating system environment variables will be loaded.")
	}
	return &Configuration{
		SERVER: Server{
			DOMAIN:      os.Getenv("SERVER_DOMAIN"),
			SERVER_PORT: os.Getenv("SERVER_PORT"),
		},
		DATABASE: Database{
			DB_HOST:     os.Getenv("DATABASE_HOST"),
			DB_PORT:     os.Getenv("DATABASE_PORT"),
			DB_USERNAME: os.Getenv("DATABASE_USERNAME"),
			DB_PASSWORD: os.Getenv("DATABASE_PASSWORD"),
			DB_NAME:     os.Getenv("DATABASE_NAME"),
		},
	}, nil
}