package interface_chi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"kaduhod/fin_v3/core/domain/investment"
	infra_investment "kaduhod/fin_v3/core/infra/investment/decimal"
	validators_dto "kaduhod/fin_v3/core/interfaces/http/dto/validators"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"time"
)
type InvestmentHandlerChi struct {
    CompoundInterestService investment.CompoundInterest
    FutureValueOfASeriesService investment.FutureValueOfASeries
}
func (h InvestmentHandlerChi) BadRequestResponse(err error, w http.ResponseWriter) {
    fmt.Println(err)
    b, err := json.Marshal(map[string]string{"_error": err.Error()})
    if err != nil {
        h.InternalServerErrorResponse(err, w)
        return
    }
    w.WriteHeader(400)
    w.Header().Add("Content-Type", "application/json")
    w.Write(b)
}
func (h InvestmentHandlerChi) BadRequestResponseWithMapOfErrors(errs map[string]string, w http.ResponseWriter) {
    fmt.Println(errs)
    w.Header().Add("Content-Type", "application/json")
    res := map[string]any{
        "_errors": errs,
    }
    b, err := json.Marshal(res)
    if err != nil {
        h.InternalServerErrorResponse(err, w)
        return
    }
    w.WriteHeader(400)
    w.Write(b)
}
func (h InvestmentHandlerChi) InternalServerErrorResponse(err error, w http.ResponseWriter) {
    fmt.Println(err)
    w.WriteHeader(http.StatusInternalServerError)
}
func (h InvestmentHandlerChi) HandleJsonBody(r *http.Request, w http.ResponseWriter) (error, []byte) {
    b, err := io.ReadAll(r.Body)
    if len(b) == 0 {
        h.BadRequestResponse(errors.New("body cannot be empty"), w)
        return nil, b
    }
    if err != nil {
        h.InternalServerErrorResponse(err, w)
        return nil, b
    }
    return nil, b
}
func (h InvestmentHandlerChi) SuccessJsonResponse(w http.ResponseWriter, content map[string]any) {
    res := map[string]any{
        "message": "ok",
        "data": content,
    }
    b, err := json.Marshal(res)
    if err != nil {
        h.InternalServerErrorResponse(err, w)
        return
    }
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(b)
}
func (h *InvestmentHandlerChi) CompoundInterestApi(w http.ResponseWriter, r *http.Request) {
    err, b := h.HandleJsonBody(r, w)
    if err != nil {
        fmt.Println(err, "asdkjaskjd")
        return
    }
    err, cpInput := struct_utils.FromJson[validators_dto.CoumpoundInterestInput](b)
    if err != nil {
        fmt.Println("Aquiii")
        h.BadRequestResponse(err, w)
        return
    }
    if err = cpInput.Validate(cpInput); err != nil {
        validations_messages := cpInput.FormatValidationError(err, "en")
        h.BadRequestResponseWithMapOfErrors(validations_messages, w)
        return
    }
    cp := h.CompoundInterestService.Calculate(
        infra_investment.NewDecimalMoney(cpInput.InitialValue),
        infra_investment.NewDecimalMoney(cpInput.TaxDecimal),
        cpInput.Periods,
    )
    res := map[string]any{
        "compound_interest": cp.GetAmount(),
        "compound_interest_money": cp.Formatted(),
        "parameters": cpInput,
    }
    h.SuccessJsonResponse(w, res)
    return
}
func (h *InvestmentHandlerChi) FutureValueOfASeriesApi(w http.ResponseWriter, r *http.Request) {
    err, b := h.HandleJsonBody(r, w)
    if err != nil {
        return
    }
    err, fvInput := struct_utils.FromJson[validators_dto.FutureValueOfASeriesInput](b)
    if err != nil {
        h.BadRequestResponse(err, w)
        return
    }
    if err := fvInput.Validate(fvInput); err != nil {
        validations_messages := fvInput.FormatValidationError(err, "en")
        h.BadRequestResponseWithMapOfErrors(validations_messages, w)
        return
    }
    fv := h.FutureValueOfASeriesService.Calculate(
        infra_investment.NewDecimalMoney(fvInput.Contribution),
        infra_investment.NewDecimalMoney(fvInput.TaxDecimal),
        fvInput.FirstDay, fvInput.Periods,
    )
    res := map[string]any{
        "future_value": fv.GetAmount(),
        "future_value_money": fv.Formatted(),
        "parameters": fvInput,
    }
    h.SuccessJsonResponse(w, res)
    return
}
func (h InvestmentHandlerChi) PredictFV(w http.ResponseWriter, r *http.Request) {
    err, b := h.HandleJsonBody(r, w)
    if err != nil {
        return
    }
    err, predictInput := struct_utils.FromJson[validators_dto.PredictContributionFVSInput](b)
    if err != nil {
        h.BadRequestResponse(err, w)
        return
    }
    if err := predictInput.Validate(predictInput); err != nil {
        validations_messages := predictInput.FormatValidationError(err, "en")
        h.BadRequestResponseWithMapOfErrors(validations_messages, w)
        return
    }
    contribution := h.FutureValueOfASeriesService.PredictContribution(
        infra_investment.NewDecimalMoney(predictInput.FinalValue),
        infra_investment.NewDecimalMoney(predictInput.TaxDecimal),
        infra_investment.NewDecimalMoney(predictInput.InitialValue),
        predictInput.ContributionOnFirstDay,
        predictInput.Periods,
    )
    res := map[string]any{
        "contribution": contribution.GetAmount(),
        "contribution_money": contribution.Formatted(),
        "parameters": predictInput,
    }
    h.SuccessJsonResponse(w, res)
    return
}
func (h *InvestmentHandlerChi) FutureValueOfASeriesWithTrackingApi(w http.ResponseWriter, r *http.Request) {
    err, b := h.HandleJsonBody(r, w)
    if err != nil {
        return
    }
    err, fvInput := struct_utils.FromJson[validators_dto.FutureValueOfASeriesWithPeriodsInput](b)
    if err != nil {
        h.BadRequestResponse(err, w)
        return
    }
    if err := fvInput.Validate(fvInput); err != nil {
        validations_messages := fvInput.FormatValidationError(err, "en")
        h.BadRequestResponseWithMapOfErrors(validations_messages, w)
        return
    }
    date_layout := "02/01/2006"
    parsedDate, _ := time.Parse(date_layout, fvInput.InitialDate)
    fv, periods := h.FutureValueOfASeriesService.CalculateTrackingPeriods(
        infra_investment.NewDecimalMoney(fvInput.InitialValue),
        infra_investment.NewDecimalMoney(fvInput.Contribution),
        infra_investment.NewDecimalMoney(fvInput.TaxDecimal),
        fvInput.FirstDay, parsedDate, fvInput.Periods,
    )
    res := map[string]any{
        "future_value": fv.GetAmount(),
        "future_value_money": fv.Formatted(),
        "periods": periods,
        "parameters": fvInput,
    }
    h.SuccessJsonResponse(w, res)
    return
}
func (h *InvestmentHandlerChi) CompoundInterest(w http.ResponseWriter, r *http.Request) {
}
func (h *InvestmentHandlerChi) FutureValueOfASeries(w http.ResponseWriter, r *http.Request) {
}
