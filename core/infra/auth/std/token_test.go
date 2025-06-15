package auth_std

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func TestToken(t *testing.T) {
    godotenv.Load()
    b, err := bcrypt.GenerateFromPassword([]byte("CtL44Rd6-zKv"), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}
