package http

type SessionService interface {
    Store(id string, data map[string]string)
    Get(id string) (map[string]string, error)
    Set(id string, key string, value string) error
}
