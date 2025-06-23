package interface_chi

import (
	"errors"
	"fmt"
	app_account_dto "kaduhod/fin_v3/core/application/account/dto"
	core_http "kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/user"
	"kaduhod/fin_v3/core/interfaces/web/renderer"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
)

type UserHandlerWeb struct {
    createUserService user.CreateUserServiceI[app_account_dto.CreateUserInput]
    signInService core_http.SigninServiceI
    renderer *renderer.Renderer
    sessionService core_http.SessionService

}
func NewUserHandlerWeb(createUserService user.CreateUserServiceI[app_account_dto.CreateUserInput], signInService core_http.SigninServiceI, sessionService core_http.SessionService, renderer *renderer.Renderer) core_http.UserHandlerWeb {
    return &UserHandlerWeb{
        createUserService: createUserService,
        signInService: signInService,
        sessionService: sessionService,
        renderer: renderer,
    }
}
func (h *UserHandlerWeb) GetSessionService() core_http.SessionService {
    return h.sessionService
}
func (h UserHandlerWeb) SignInForm(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    data := map[string]any{
        "csrf": csrf,
    }
    if err := h.renderer.Render(w, "signin_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h UserHandlerWeb) SignUpForm(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    data := map[string]any{
        "csrf": csrf,
    }
    if err := h.renderer.Render(w, "signup_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h UserHandlerWeb) SignIn(w http.ResponseWriter, r *http.Request) {
}
func (h UserHandlerWeb) SignUp(w http.ResponseWriter, r *http.Request) {
}
func (h *UserHandlerWeb) getSession(r *http.Request) (core_http.SessionData, error) {
    session, err := h.sessionService.Get(struct_utils.GetCookie(r).Value)
    if err != nil {
        fmt.Println(err)
        return session, err
    }
    return session, nil
}
func (h *UserHandlerWeb) getCsrfToken(r *http.Request) (string, error) {
    session, err := h.getSession(r)
    if err != nil {
        return "", errors.New("Session not found")
    }
    return session.Csrf, nil
}
