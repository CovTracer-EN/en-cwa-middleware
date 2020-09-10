package main

import (
	"log"

	"github.com/rise-center/en-cwa-middleware/internal/http"
)

func main() {
	if err := http.Run(); err != nil {
		log.Fatal(err)
	}
}
