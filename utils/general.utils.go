package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string{
	godotenv.Load(".env")

	value := os.Getenv(key)

	return value
}