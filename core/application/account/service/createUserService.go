package app_account_service

import (
	"errors"
	"fmt"
	app_account_dto "kaduhod/fin_v3/core/application/account/dto"
	"kaduhod/fin_v3/core/domain/repository"
	"kaduhod/fin_v3/core/domain/user"

	"golang.org/x/crypto/bcrypt"
)
type CreateUserService struct {
    repository repository.Repository[user.User]
}
func NewCreateUserService(repository repository.Repository[user.User]) user.CreateUserServiceI[app_account_dto.CreateUserInput] {
    return &CreateUserService{
        repository: repository,
    }
}
func (s *CreateUserService) Create(usr app_account_dto.CreateUserInput) error {
    fmt.Println(usr.Password)
    hash , err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println(err)
        return err
    }
    dbUser, err := s.repository.Get(user.User{Email: usr.Email})
    if err != nil && err.Error() != "failed to get user: no rows in result set" {
        return err
    }
    if dbUser.Email != "" {
        fmt.Println(err)
        return errors.New("User email not available")
    }
    if _, err := s.repository.Save(user.User{
        Name:     usr.Name,
        Password: string(hash),
        Email:    usr.Email,
    }); err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}
