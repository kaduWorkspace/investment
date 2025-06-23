package http

import "kaduhod/fin_v3/core/domain/user"

type SigninServiceI interface {
    Signin(usr user.User, password string) (error)
}
