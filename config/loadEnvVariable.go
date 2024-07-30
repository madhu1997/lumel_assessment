package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load("app_env.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
