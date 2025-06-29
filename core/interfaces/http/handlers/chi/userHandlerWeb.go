package interface_chi

import (
	"errors"
	"fmt"
	app_account_dto "kaduhod/fin_v3/core/application/account/dto"
	core_http "kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/repository"
	"kaduhod/fin_v3/core/domain/user"
	validators_dto "kaduhod/fin_v3/core/interfaces/http/dto/validators"
	"kaduhod/fin_v3/core/interfaces/web/renderer"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"strings"
)

type UserHandlerWeb struct {
    userRepo repository.Repository[user.User]
    createUserService user.CreateUserServiceI[app_account_dto.CreateUserInput]
    signInService core_http.SigninServiceI
    renderer *renderer.Renderer
    sessionService core_http.SessionService
}
func NewUserHandlerWeb(userRepo repository.Repository[user.User] ,createUserService user.CreateUserServiceI[app_account_dto.CreateUserInput], signInService core_http.SigninServiceI, sessionService core_http.SessionService, renderer *renderer.Renderer) core_http.UserHandlerWeb {
    return &UserHandlerWeb{
        userRepo: userRepo,
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
func (h UserHandlerWeb) SignUp(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    data := map[string]any{
        "csrf": csrf,
    }
    userInput := validators_dto.NewCreateUserInput(
        r.FormValue("email"),
        r.FormValue("name"),
        r.FormValue("password"),
        r.FormValue("password_confirm"),
    ).(validators_dto.CreateUserInput)
    if err := userInput.Validate(); err != nil {
        errorMessages := userInput.FormatValidationError(err, "pt")
        data["errs"] = errorMessages
        if err := h.renderer.Render(w, "signup_page", data); err != nil {
            fmt.Println(err)
        }
        return
    }
    createUserInput := app_account_dto.NewCreateUserInput(userInput.Name, userInput.Email, userInput.Password).(app_account_dto.CreateUserInput)
    if err := h.createUserService.Create(createUserInput); err != nil {
        if err := h.renderer.Render(w, "signup_page", data); err != nil {
            fmt.Println(err)
            w.WriteHeader(http.StatusBadRequest)
        } else {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }
    data["message"] = "success"
    if err := h.renderer.Render(w, "signin_page", data); err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Println(err)
    }
}
func (h UserHandlerWeb) SignIn(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    data := map[string]any{
        "csrf": csrf,
    }
    email := strings.ToLower(r.FormValue("email"))
    userInput := validators_dto.NewSignInInput(
        email,
        r.FormValue("password"),
    )
    if err := userInput.Validate(); err != nil {
        errorMessages := userInput.FormatValidationError(err, "pt")
        data["errs"] = errorMessages
        fmt.Println(errorMessages)
        if err := h.renderer.Render(w, "signin_page", data); err != nil {
            fmt.Println(err)
        }
        return
    }
    if err := h.signInService.Signin(user.User{
        Email: email,
    }, r.FormValue("password")); err != nil {
        fmt.Println(err)
        if err.Error() == "User not found" {
            data["errs"] = []string{"Account not found"}
            if err := h.renderer.Render(w, "signin_page", data); err != nil {
                fmt.Println(err)
            }
        } else if err.Error() == "Invalid password" {
            data["errs"] = []string{"Invalid password"}
            if err := h.renderer.Render(w, "signin_page", data); err != nil {
                fmt.Println(err)
            }
        } else {
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }
    u, err := h.userRepo.Get(user.User{Email: email})
    if err != nil {
        fmt.Println(err)
        if err := h.renderer.Render(w, "home", data); err != nil {
            fmt.Println(err)
        }
        return
    }
    session, err := h.getSession(r)
    if err != nil {
        fmt.Println(err)
        if err := h.renderer.Render(w, "home", data); err != nil {
            fmt.Println(err)
        }
        return
    }
    session.Usr = u
    h.sessionService.Store(struct_utils.GetCookie(r).Value, session)
    w.Header().Set("HX-Redirect", "/web/dashboard")
    http.Redirect(w, r, "/web/dashboard", http.StatusSeeOther)
}
func (h *UserHandlerWeb) SignOut(w http.ResponseWriter, r *http.Request) {
    if err := h.sessionService.Destroy(struct_utils.GetCookie(r).Value); err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.Header().Set("HX-Redirect", "/")
    http.Redirect(w, r, "/", http.StatusSeeOther)
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
