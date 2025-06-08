package memory

import (
	"errors"
	"kaduhod/fin_v3/core/domain/http"
	"sync"
)

type InMemorySession struct {
    data map[string]map[string]string
    mu *sync.RWMutex
}
var (
    data = make(map[string]map[string]string)
    mu  = &sync.RWMutex{}
)
func NewInMemorySession() http.SessionService {
    return &InMemorySession{
        data: data,
        mu: mu,
    }
}
func (s *InMemorySession) Get(id string) (map[string]string, error) {
    s.mu.Lock()
    valor, ok := s.data[id]
    s.mu.Unlock()
    if !ok {
        return nil, errors.New("Id not found")
    }
    return valor, nil
}
func (s *InMemorySession) Store(id string, sessionData map[string]string) {
    s.mu.Lock()
    s.data[id] = sessionData
    s.mu.Unlock()
}
func (s *InMemorySession) Set(id string, key string, value string) error {
    s.mu.Lock()
    _, ok := s.data[id]
    if !ok {
        s.mu.Unlock()
        return errors.New("Id not found")
    }
    s.data[id][key] = value
    s.mu.Unlock()
    return nil
}
