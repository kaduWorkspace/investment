package http

type SessionData struct {
    Expiration int64
    Csrf       string
}
type SessionService interface {
    Store(id string, data SessionData)
    Get(id string) (SessionData, error)
}
