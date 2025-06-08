package http_middleware

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	core_http "kaduhod/fin_v3/core/domain/http"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"os"
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
func (m *SessionHandlerMiddleware) createSession(w http.ResponseWriter, r *http.Request) core_http.SessionData {
    cookie := struct_utils.GetCookie(r)
    var sessionData core_http.SessionData
    if cookie == nil {
        cookie = struct_utils.CreateCookie(w)
    }
    csrf, err := m.createCsrfToken()
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return sessionData
    }
    sessionData = core_http.SessionData {
        Expiration: cookie.Expires.Unix(),
        Csrf: csrf,
    }
    m.sessionService.Store(cookie.Value, sessionData)
    return sessionData
}
func (m *SessionHandlerMiddleware) validateSession(session core_http.SessionData) bool {
    expirationDate := time.Unix(session.Expiration, 0)
    if expirationDate.Before(time.Now()) {
        return false
    }
    return true
}
func (m *SessionHandlerMiddleware) getSession(r *http.Request) (core_http.SessionData, error) {
    var session core_http.SessionData
    if r == nil {
        return session, errors.New("Request is nil")
    }
    cookie := struct_utils.GetCookie(r)
    if cookie == nil {
        return session, errors.New("Cookie is nil")
    }
    session, err := m.sessionService.Get(cookie.Value)
    if err != nil {
        fmt.Println(err)
        return session, err
    }
    return session, nil
}
func (m *SessionHandlerMiddleware) CreateSessionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := m.getSession(r)
        if err != nil && (err.Error() != "Id not found" && err.Error() != "Cookie is nil"){
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        if m.validateSession(session) {
            next.ServeHTTP(w, r)
            return
        }
        _ = m.createSession(w, r)
        next.ServeHTTP(w, r)
    })
}
func (m *SessionHandlerMiddleware) CheckSessionMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := m.getSession(r);
        if err != nil && (err.Error() != "Cookie is nil" && err.Error() != "Id not found"){
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        if m.validateSession(session) {
            next.ServeHTTP(w, r)
            return
        }
        w.Header().Set("HX-Redirect", "/")
        w.WriteHeader(303)
        http.Redirect(w, r, "/", http.StatusSeeOther)
    })
}
func (m *SessionHandlerMiddleware) createCsrfToken() (string, error) {
    nonce := make([]byte, 32)
    if _, err := rand.Read(nonce); err != nil {
        return "", err
    }
    signed, err := m.sign(nonce)
    if err != nil {
        return "", err
    }
    return signed, nil
}
func (m *SessionHandlerMiddleware) sign(nonce []byte) (string, error){
    secret := os.Getenv("CSRF_SECRET")
    h := hmac.New(sha256.New, []byte(secret))
    h.Write([]byte(nonce))
    signed := h.Sum(nil)
    token := base64.StdEncoding.EncodeToString(nonce) + "." + base64.StdEncoding.EncodeToString(signed)
    return token, nil
}
