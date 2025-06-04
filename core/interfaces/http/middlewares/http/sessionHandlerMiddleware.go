package http_middleware

import (
	"fmt"
	core_http "kaduhod/fin_v3/core/domain/http"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"strconv"
	"time"
)

type SessionHandlerMiddleware struct {
    sessionService core_http.SessionService
}
func NewSessionHandlerMiddleware(sessionService core_http.SessionService) *SessionHandlerMiddleware {
    return &SessionHandlerMiddleware{
        sessionService: sessionService,
    }
}
func (m *SessionHandlerMiddleware) createSession(sessionId string) map[string]string {
    expirationTime := fmt.Sprintf("%d",time.Now().Add(72 * time.Hour).Unix())
    sessionData := map[string]string{
        "expiration": expirationTime,
    }
    m.sessionService.Store(sessionId, sessionData)
    fmt.Println("Creating session", sessionData)
    return sessionData
}
func (m *SessionHandlerMiddleware) validateSession(session map[string]string) bool {
    expirationStr, ok := session["expiration"]
    if !ok {
        return false
    }
    expiration, err := strconv.Atoi(expirationStr)
    if err != nil {
        fmt.Println(err)
        return false
    }
    expirationDate := time.Unix(int64(expiration), 0)
    if expirationDate.Before(time.Now()) {
        fmt.Println("Session expired")
        return false
    }
    fmt.Println("Session validated", session)
    return true
}
func (m *SessionHandlerMiddleware) getSession(r *http.Request) map[string]string {
    sessionId := struct_utils.SessionId(r)
    session, err := m.sessionService.Get(sessionId)
    if err != nil {
        fmt.Println(err)
        return nil
    }
    return session
}
func (m *SessionHandlerMiddleware) CreateSessionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var session map[string]string
        if session = m.getSession(r); session != nil && m.validateSession(session) {
            next.ServeHTTP(w, r)
            return
        }
        _ = m.createSession(struct_utils.SessionId(r))
        next.ServeHTTP(w, r)
    })
}
func (m *SessionHandlerMiddleware) CheckSessionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var session map[string]string
        if session = m.getSession(r); session != nil && m.validateSession(session) {
            next.ServeHTTP(w, r)
            return
        }
        w.WriteHeader(302)
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })
}
