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

var AppEnv Env

func Init() error {
	if err := loadEnv(); err != nil {
		return err
	}

	setEnv()

	return nil
}

func loadEnv() error {

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	_ = godotenv.Load(".env." + env + ".local")
	if "test" != env {
		_ = godotenv.Load(".env.local")
	}
	_ = godotenv.Load(".env." + env)
	_ = godotenv.Load() // The Original .env

	return nil
}

func setEnv() {
	switch os.Getenv("APP_ENV") {
	case "development":
		AppEnv = Development
	case "development.docker":
		AppEnv = DevelopmentDocker
	default:
		AppEnv = Production
	}
}
