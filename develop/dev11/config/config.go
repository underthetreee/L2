package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP HTTP
}

type HTTP struct {
	Port int
}

func NewConfig() (*Config, error) {
	err := godotenv.Load(".env")

	portEnv := os.Getenv("HTTP_PORT")
	port, err := strconv.Atoi(portEnv)
	if err != nil {
		return nil, err
	}

	return &Config{
		HTTP: HTTP{
			Port: port,
		},
	}, nil
}
