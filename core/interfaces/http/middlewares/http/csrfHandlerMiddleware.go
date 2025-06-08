package http_middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	core_http "kaduhod/fin_v3/core/domain/http"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"os"
	"strings"
)

type CsrfHandlerMiddleware struct {
    sessionService core_http.SessionService
}
func NewCsrfHandlerMiddleware(sessionService core_http.SessionService) *CsrfHandlerMiddleware {
    return &CsrfHandlerMiddleware{
        sessionService: sessionService,
    }
}
func (m *CsrfHandlerMiddleware) getSession(r *http.Request) (core_http.SessionData, error) {
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
func (m *CsrfHandlerMiddleware) ValidateCsrfMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        csrf := r.FormValue("_csrf")
        if csrf == "" {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        if !m.validateCsrfToken(csrf) {
            w.WriteHeader(http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}
func (m *CsrfHandlerMiddleware) validateCsrfToken(token string) bool {
    parts := strings.Split(token, ".")
    if len(parts) != 2 {
        fmt.Println("Not enough parts", parts)
        return false
    }
    nonce, err := base64.StdEncoding.DecodeString(parts[0])
    signReceived, err2 := base64.StdEncoding.DecodeString(parts[1])
    if err != nil || err2 != nil {
        fmt.Println(err)
        return false
    }
    mac := hmac.New(sha256.New, []byte(os.Getenv("CSRF_SECRET")))
    mac.Write(nonce)
    expectedSign := mac.Sum(nil)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return hmac.Equal([]byte(expectedSign), []byte(signReceived))
}
