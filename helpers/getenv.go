package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func GetEnvData() Env {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVar := Env{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}

	return envVar
}
