package interface_chi

import (
	"errors"
	"fmt"
	"kaduhod/fin_v3/core/domain/external"
	core_http "kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/investment"
	"kaduhod/fin_v3/core/domain/repository"
	"kaduhod/fin_v3/core/domain/user"
	"kaduhod/fin_v3/core/interfaces/web/renderer"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
)
type DashboardHandlerWeb struct {
    bcbService external.BcbI
    userRepo repository.Repository[user.User]
    renderer *renderer.Renderer
    sessionService core_http.SessionService
    FutureValueOfASeriesService investment.FutureValueOfASeries
}
func (h *DashboardHandlerWeb) initData(r *http.Request) (map[string]any, error) {
    session, err := h.getSession(r)
    if err != nil {
        return nil, err
    }
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        return nil, err
    }
    return map[string]any{
        "logged": true,
        "user": session.Usr,
        "csrf": csrf,
    }, nil
}
func NewDashboardHandlerWeb(futureValueOfASeriesService investment.FutureValueOfASeries ,bcb external.BcbI, userRepo repository.Repository[user.User], sessionService core_http.SessionService, renderer *renderer.Renderer) core_http.DashboardHandler {
    return &DashboardHandlerWeb{
        bcbService: bcb,
        FutureValueOfASeriesService: futureValueOfASeriesService,
        userRepo: userRepo,
        sessionService: sessionService,
        renderer: renderer,
    }
}
func (h *DashboardHandlerWeb) FVSDashboard(w http.ResponseWriter, r *http.Request) {
    data, err := h.initData(r)
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := h.renderer.Render(w, "dashboard_fv", data); err != nil {
        fmt.Println(err)
    }
}
func (h *DashboardHandlerWeb) PredictDashboard(w http.ResponseWriter, r *http.Request) {
    data, err := h.initData(r)
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := h.renderer.Render(w, "dashboard_predict", data); err != nil {
        fmt.Println(err)
    }
}
func (h *DashboardHandlerWeb) FVS(w http.ResponseWriter, r *http.Request) {

}
func (h *DashboardHandlerWeb) Predict(w http.ResponseWriter, r *http.Request) {

}
func (h *DashboardHandlerWeb) Index(w http.ResponseWriter, r *http.Request) {
    data, err := h.initData(r)
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := h.renderer.Render(w, "dashboard", data); err != nil {
        fmt.Println(err)
    }
}
func (h *DashboardHandlerWeb) Dashboard(w http.ResponseWriter, r *http.Request) {
    data, err := h.initData(r)
    if err != nil {
        fmt.Println(err)
        return
    }
    if err := h.renderer.Render(w, "dashboard_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h *DashboardHandlerWeb) getSession(r *http.Request) (core_http.SessionData, error) {
    session, err := h.sessionService.Get(struct_utils.GetCookie(r).Value)
    if err != nil {
        fmt.Println(err)
        return session, err
    }
    return session, nil
}
func (h *DashboardHandlerWeb) getCsrfToken(r *http.Request) (string, error) {
    session, err := h.getSession(r)
    if err != nil {
        return "", errors.New("Session not found")
    }
    return session.Csrf, nil
}
func (h *DashboardHandlerWeb) GetSessionService() core_http.SessionService {
    return h.sessionService
}
