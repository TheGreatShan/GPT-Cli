package services

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetKey(name string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	return os.Getenv(name)
}
