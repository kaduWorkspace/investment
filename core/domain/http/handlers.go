package http
import "net/http"

type InvestmentHandler interface {
    CompoundInterestApi(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesApi(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesWithTrackingApi(w http.ResponseWriter, r *http.Request)
    PredictFV(w http.ResponseWriter, r *http.Request)
}

type InvestmentHandlerWeb interface {
    FutureValueOfASeriesPredictFormPage(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesFormPage(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesResultPage(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesPredictResultPage(w http.ResponseWriter, r *http.Request)
}
