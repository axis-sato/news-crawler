package utils

import (
	"github.com/joho/godotenv"
	"os"
)

type Env int

const (
	Development Env = iota
	Production
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return  err
	}
	return nil
}

func GetEnv() Env {
	switch os.Getenv("APP_ENV") {
	case "development":
		return Development
	case "production":
		return Production
	default:
		return Development
	}
}
