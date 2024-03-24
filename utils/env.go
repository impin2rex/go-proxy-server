package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type env struct {
	API_URL string
	API_KEY string
}

var ENV env

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		panic("Cannot load env")
	}
	ENV = env{
		API_URL: os.Getenv("API_URL"),
		API_KEY: os.Getenv("API_KEY"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
