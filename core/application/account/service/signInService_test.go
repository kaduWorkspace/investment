package app_account_service

import (
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	"os"
	"testing"

	"github.com/joho/godotenv"
)
func setupDB(t *testing.T) *pg_connection.PgxConnextion {
	t.Helper()
    enviroment := os.Args[len(os.Args) - 1]
    envFile := ".env.development"
    if enviroment == "local" {
        envFile = ".env.local"
    }
    err := godotenv.Load("../../../../" + envFile)
    if err != nil {
        t.Error(err)
        t.Fail()
    }

	conn := pg_connection.NewPgxConnection()
	return conn
}
