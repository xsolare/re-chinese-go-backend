package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


var POSTGRES_HOST string
var POSTGRES_PORT string
var POSTGRES_USER string
var POSTGRES_PASSWORD string
var POSTGRES_DB string

func GetEnvVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
}
