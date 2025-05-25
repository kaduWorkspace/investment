package main

import (
	main_chi "kaduhod/fin_v3/cmd/web/chi"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    main_chi.Run()
}
