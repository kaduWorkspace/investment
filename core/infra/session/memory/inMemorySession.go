package memory

import (
	"errors"
	"kaduhod/fin_v3/core/domain/http"
	"sync"
)

type InMemorySession struct {
    data map[string]http.SessionData
    mu *sync.RWMutex
}
var (
    data = make(map[string]http.SessionData)
    mu  = &sync.RWMutex{}
)
func NewInMemorySession() http.SessionService {
    return &InMemorySession{
        data: data,
        mu: mu,
    }
}
func (s *InMemorySession) Get(id string) (http.SessionData, error) {
    s.mu.Lock()
    var sessionData http.SessionData
    sessionData, ok := s.data[id]
    s.mu.Unlock()
    if !ok {
        return sessionData, errors.New("Id not found")
    }
    return sessionData, nil
}
func (s *InMemorySession) Store(id string, sessionData http.SessionData) {
    s.mu.Lock()
    s.data[id] = sessionData
    s.mu.Unlock()
}
func (s *InMemorySession) Destroy(id string) error {
    s.mu.Lock()
    delete(s.data, id)
    s.mu.Unlock()
    return nil
}
