package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key, def string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	res := os.Getenv(key)
	if res != "" {
		return res
	}
	return def
}
