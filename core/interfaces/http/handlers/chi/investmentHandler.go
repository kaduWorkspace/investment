package interface_chi

import (
	"encoding/json"
	"fmt"
	"io"
	"kaduhod/fin_v3/core/domain/investment"
	validators_dto "kaduhod/fin_v3/core/interfaces/http/dto/validators"
	"net/http"
)
type InvestmentHandlerChi struct {
    CompoundInterestService investment.CompoundInterest
    FutureValueOfASeriesService investment.FutureValueOfASeries
}
func (h InvestmentHandlerChi) InternalServerErrorResponse(err error, w http.ResponseWriter) {
    fmt.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
}
func (h *InvestmentHandlerChi) CompoundInterestApi(w http.ResponseWriter, r *http.Request) {
    b, err := io.ReadAll(r.Body)
    w.Header().Add("Content-Type", "application/json")
    if len(b) == 0 {
        w.WriteHeader(400)
        w.Write([]byte(`{"_error":"body cannot be empty"}`))
        return
    }
    if err != nil {
        h.InternalServerErrorResponse(err, w)
        return
    }
    var body validators_dto.CoumpoundInterestInput
    if err = json.Unmarshal(b, &body); err != nil {
        h.InternalServerErrorResponse(err, w)
        return
    }
    err = body.Validate(body)
    if err != nil {
        messages := body.FormatValidationError(err, "en")
        b, err := json.Marshal(messages)
        if err != nil {
            h.InternalServerErrorResponse(err, w)
            return
        }
        w.WriteHeader(400)
        w.Write(b)
        return
    }
    w.WriteHeader(200)
    return
}
func (h *InvestmentHandlerChi) FutureValueOfASeriesApi(w http.ResponseWriter, r *http.Request) {
}
func (h *InvestmentHandlerChi) CompoundInterest(w http.ResponseWriter, r *http.Request) {
}
func (h *InvestmentHandlerChi) FutureValueOfASeries(w http.ResponseWriter, r *http.Request) {
}
