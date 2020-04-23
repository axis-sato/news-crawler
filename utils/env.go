package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Env int

const (
	Development Env = iota
	DevelopmentDocker
	Production
)

func LoadEnv() error {

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	_ = godotenv.Load(".env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load()

	return nil
}

func GetEnv() Env {
	switch os.Getenv("APP_ENV") {
	case "development":
		return Development
	case "development-docker":
		return DevelopmentDocker
	case "production":
		return Production
	default:
		return Development
	}
}
