package auth_std

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func TestToken(t *testing.T) {
    godotenv.Load()
    b, err := bcrypt.GenerateFromPassword([]byte("yp7ORf71nOKza4yGab5O7x1p"), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}
