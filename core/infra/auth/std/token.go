package auth_std

import (
	"fmt"
	core_http "kaduhod/fin_v3/core/domain/http"
	"net/http"
	"os"

	"golang.org/x/crypto/bcrypt"
)
/**************************************************
Autenticação por header de request e token de admin
**************************************************/
type IdentityToken struct {
    request *http.Request
}
func (u IdentityToken) GetIndentidy() (interface{}, error) {
    return u.request.Header.Get("Authorization"), nil
}
type AuthToken struct {}
func newAuthService() core_http.AuthService {
    return AuthToken{}
}
func (a AuthToken) Authenticate(user core_http.Identity) (bool, error) {
    headerToken, err := user.GetIndentidy()
    if err != nil {
        fmt.Println(err)
        return false, err
    }
    fmt.Println(os.Getenv("APP_ADMIN_HASH"))
    if err := bcrypt.CompareHashAndPassword([]byte(os.Getenv("APP_ADMIN_HASH")), []byte(headerToken.(string))); err != nil {
        fmt.Println(err)
        return false, err
    }
    return true, nil
}
func AuthTokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        identity := IdentityToken{request: r}
        authService := newAuthService()
        auth, err := authService.Authenticate(identity)
        if auth && err == nil {
            next.ServeHTTP(w, r)
        } else {
            w.WriteHeader(401)
        }
    })
}
