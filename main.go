package main

import (
	"log"
	"os"

	"github.com/rise-center/en-cwa-middleware/internal/http"

	"github.com/joho/godotenv"
)

func main() {
	envPath := os.Getenv("ENV")
	if envPath == "" {
		envPath = ".env"
	}
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	if err := http.Run(); err != nil {
		log.Fatal(err)
	}
}
