package auth_std

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestToken(t *testing.T) {
    b, err := bcrypt.GenerateFromPassword([]byte("XtKw4l\v07m'"), bcrypt.DefaultCost)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))
}
