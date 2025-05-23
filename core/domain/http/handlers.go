package http
import "net/http"

type InvestmentHandler interface {
    CompoundInterestApi(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeriesApi(w http.ResponseWriter, r *http.Request)
    CompoundInterest(w http.ResponseWriter, r *http.Request)
    FutureValueOfASeries(w http.ResponseWriter, r *http.Request)
}
