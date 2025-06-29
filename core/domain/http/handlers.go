package http
import "net/http"

type Handler interface {
    GetSessionService() SessionService
}

type InvestmentHandler interface {
    CompoundInterestApi(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesApi(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesWithTrackingApi(w http.ResponseWriter, r *http.Request)
    PredictFV(w http.ResponseWriter, r *http.Request)
}

type InvestmentHandlerWeb interface {
    Handler
    Index(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesPredictFormPage(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesFormPage(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesResultPage(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesPredictResultPage(w http.ResponseWriter, r *http.Request)
}

type UserHandlerWeb interface {
    Handler
    SignInForm(w http.ResponseWriter, r *http.Request)
    SignUpForm(w http.ResponseWriter, r *http.Request)
    SignIn(w http.ResponseWriter, r *http.Request)
    SignUp(w http.ResponseWriter, r *http.Request)
    SignOut(w http.ResponseWriter, r *http.Request)
}

type DashboardHandler interface {
    Handler
    Dashboard(w http.ResponseWriter, r *http.Request)
    Index(w http.ResponseWriter, r *http.Request)
    FVSDashboard(w http.ResponseWriter, r *http.Request)
    PredictDashboard(w http.ResponseWriter, r *http.Request)
    FVS(w http.ResponseWriter, r *http.Request)
    Predict(w http.ResponseWriter, r *http.Request)
}
