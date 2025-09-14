package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetEnv(key string, fallback string) string {
	env, present := os.LookupEnv(key)

	if !present {
		return fallback
	}

	return env
}
