package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(pathFile string) {
	err := godotenv.Load(pathFile)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
