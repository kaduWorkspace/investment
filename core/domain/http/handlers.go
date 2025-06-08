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
