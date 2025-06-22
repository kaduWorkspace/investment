package main

import (
	main_chi "kaduhod/fin_v3/cmd/web/chi"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    envFile := ""
    if os.Args[1] == "local" {
        envFile = ".local"
    }
    if os.Args[1] == "development" {
        envFile = ".development"
    }
    err := godotenv.Load("./.env" + envFile)
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    main_chi.Run()
}
