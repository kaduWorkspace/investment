package http

import "kaduhod/fin_v3/core/domain/user"

type SessionData struct {
    Expiration int64
    Csrf       string
    Usr        user.User
}
type SessionService interface {
    Store(id string, data SessionData)
    Get(id string) (SessionData, error)
}
