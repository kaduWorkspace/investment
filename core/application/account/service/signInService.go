package app_account_service

import (
	"errors"
	"fmt"
	"kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/repository"
	"kaduhod/fin_v3/core/domain/user"

	"golang.org/x/crypto/bcrypt"
)

type SigninService struct {
    repository repository.Repository[user.User]
}
func NewSigninService(repository repository.Repository[user.User]) http.SigninServiceI {
    return &SigninService{repository: repository}
}
func (s SigninService) Signin(usr user.User, password string) (error) {
    dbUser, err := s.repository.Get(user.User{Email: usr.Email})
    if err != nil && err.Error() != "failed to get user: no rows in result set" {
        fmt.Println(err)
        return err
    }
    if dbUser.Email != usr.Email {
        return errors.New("User not found")
    }
    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
        return errors.New("Invalid password")
    }
    return nil
}
