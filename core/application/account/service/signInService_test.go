package app_account_service

import (
	"kaduhod/fin_v3/core/domain/user"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	pg_repository "kaduhod/fin_v3/core/infra/persistence/postgres/repository"
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

func TestSignInService_Integration(t *testing.T) {
    conn := setupDB(t)
    userReposiotry := pg_repository.NewUserRepository(conn)
    signInService := NewSigninService(userReposiotry)
    t.Run("Test succesefull signin", func(t *testing.T) {
        pwd := os.Getenv("APP_ADMIN_TOKEN")
        if err := signInService.Signin(user.User{Email: "admin@admin.com"}, pwd); err != nil {
            t.Log(err)
            t.Fail()
        }
    })
    t.Run("Test failed signin", func(t *testing.T) {
        err := signInService.Signin(user.User{Email: "admin@admin.com"}, "wrong_password")
        if err == nil {
            t.Log(err)
            t.Fail()
        }
        if err.Error() != "Invalid password" {
            t.Log(err)
            t.Fail()
        }
    })
}
