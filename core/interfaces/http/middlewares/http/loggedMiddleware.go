package http_middleware

import (
	"errors"
	"fmt"
	core_http "kaduhod/fin_v3/core/domain/http"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
)
type LoggedHandlerMiddleware struct {
    sessionService core_http.SessionService
}
func NewLoggedHandlerMiddleware(sessionService core_http.SessionService) *LoggedHandlerMiddleware {
    return &LoggedHandlerMiddleware{
        sessionService: sessionService,
    }
}
func (m *LoggedHandlerMiddleware) getSession(r *http.Request) (core_http.SessionData, error) {
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
func (m *LoggedHandlerMiddleware) ValidateSession(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, err := m.getSession(r)
        if err != nil || session.Usr.Id == 0 {
            w.Header().Set("HX-Redirect", "/")
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }
        next.ServeHTTP(w, r)
    })
}
