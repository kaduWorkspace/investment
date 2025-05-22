package interface_chi

import (
	"kaduhod/fin_v3/core/domain/investment"
	"net/http"
)
type InvestmentHandlerChi struct {
    CompoundInterestService investment.CompoundInterest
    FutureValueOfASeriesService investment.FutureValueOfASeries
}
func (h *InvestmentHandlerChi) CompoundInterestApi(w http.ResponseWriter, r *http.Request) {
}
func (h *InvestmentHandlerChi) FutureValueOfASeriesApi(w http.ResponseWriter, r *http.Request) {
}
func (h *InvestmentHandlerChi) CompoundInterest(w http.ResponseWriter, r *http.Request) {
}
func (h *InvestmentHandlerChi) FutureValueOfASeries(w http.ResponseWriter, r *http.Request) {
}
