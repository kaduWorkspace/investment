package http

type Identity interface {
    GetIndentidy() (interface{}, error)
}

type AuthService interface {
    Authenticate(user Identity) (bool, error)
}
