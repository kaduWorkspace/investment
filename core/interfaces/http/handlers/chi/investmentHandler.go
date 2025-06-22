package interface_chi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	app_investment_decimal "kaduhod/fin_v3/core/application/investment/service/decimal"
	"kaduhod/fin_v3/core/domain/external"
	core_http "kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/investment"
	validators_dto "kaduhod/fin_v3/core/interfaces/http/dto/validators"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"strings"
	"time"
)
type InvestmentHandlerChi struct {
    bcbService external.BcbI
    CompoundInterestService investment.CompoundInterest
    FutureValueOfASeriesService investment.FutureValueOfASeries
}
func NewInvestmentHandler(bcbService external.BcbI, compoundInterestService investment.CompoundInterest, futureValueOfASeriesService investment.FutureValueOfASeries) core_http.InvestmentHandler {
    return &InvestmentHandlerChi{
        CompoundInterestService: compoundInterestService,
        FutureValueOfASeriesService: futureValueOfASeriesService,
        bcbService: bcbService,
    }
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
        h.BadRequestResponse(err, w)
        return
    }
    if err = cpInput.Validate(cpInput); err != nil {
        validations_messages := cpInput.FormatValidationError(err, "en")
        h.BadRequestResponseWithMapOfErrors(validations_messages, w)
        return
    }
    cp := h.CompoundInterestService.Calculate(
        app_investment_decimal.NewDecimalMoney(cpInput.InitialValue),
        app_investment_decimal.NewDecimalMoney(cpInput.TaxDecimal),
        cpInput.Periods,
    )
    cpReal := h.CompoundInterestService.CalculateRealValue(
        app_investment_decimal.NewDecimalMoney(cpInput.InitialValue),
        app_investment_decimal.NewDecimalMoney(cpInput.TaxDecimal),
        app_investment_decimal.NewDecimalMoney(cpInput.TaxDecimalInflation),
        cpInput.Periods,
    )
    res := map[string]any{
        "compound_interest": cp.GetAmount(),
        "compound_interest_money": cp.Formatted(),
        "compound_interest_real": cpReal.GetAmount(),
        "compound_interest_real_money": cpReal.Formatted(),
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
        app_investment_decimal.NewDecimalMoney(fvInput.Contribution),
        app_investment_decimal.NewDecimalMoney(fvInput.TaxDecimal),
        fvInput.FirstDay, fvInput.Periods,
    )
    fvReal := h.FutureValueOfASeriesService.CalculateRealValue(
        app_investment_decimal.NewDecimalMoney(fvInput.Contribution),
        app_investment_decimal.NewDecimalMoney(fvInput.TaxDecimal),
        app_investment_decimal.NewDecimalMoney(fvInput.TaxDecimalInflation),
        fvInput.FirstDay, fvInput.Periods,
    )
    res := map[string]any{
        "future_value": fv.GetAmount(),
        "future_value_real": fvReal.GetAmount(),
        "future_value_money": fv.Formatted(),
        "future_value_real_money": fvReal.Formatted(),
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
        app_investment_decimal.NewDecimalMoney(predictInput.FinalValue),
        app_investment_decimal.NewDecimalMoney(predictInput.TaxDecimal),
        app_investment_decimal.NewDecimalMoney(predictInput.InitialValue),
        predictInput.ContributionOnFirstDay,
        predictInput.Periods,
    )
    contributionReal := h.FutureValueOfASeriesService.PredictContributionRealValue(
        app_investment_decimal.NewDecimalMoney(predictInput.FinalValue),
        app_investment_decimal.NewDecimalMoney(predictInput.TaxDecimal),
        app_investment_decimal.NewDecimalMoney(predictInput.TaxDecimalInflation),
        app_investment_decimal.NewDecimalMoney(predictInput.InitialValue),
        predictInput.ContributionOnFirstDay,
        predictInput.Periods,
    )
    res := map[string]any{
        "contribution": contribution.GetAmount(),
        "contribution_real": contributionReal.GetAmount(),
        "contribution_money": contribution.Formatted(),
        "contribution_real_money": contributionReal.Formatted(),
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
        app_investment_decimal.NewDecimalMoney(fvInput.InitialValue),
        app_investment_decimal.NewDecimalMoney(fvInput.Contribution),
        app_investment_decimal.NewDecimalMoney(fvInput.TaxDecimal),
        fvInput.FirstDay, parsedDate, fvInput.Periods,
    )
    fvReal, periodsReal := h.FutureValueOfASeriesService.CalculateTrackingPeriodsRealValue(
        app_investment_decimal.NewDecimalMoney(fvInput.InitialValue),
        app_investment_decimal.NewDecimalMoney(fvInput.Contribution),
        app_investment_decimal.NewDecimalMoney(fvInput.TaxDecimal),
        app_investment_decimal.NewDecimalMoney(fvInput.TaxDecimalInflation),
        fvInput.FirstDay, parsedDate, fvInput.Periods,
    )

    res := map[string]any{
        "future_value": fv.GetAmount(),
        "future_value_money": fv.Formatted(),
        "periods": periods,
        "future_value_real": fvReal.GetAmount(),
        "future_value_money_real": fvReal.Formatted(),
        "periods_real": periodsReal,
        "parameters": fvInput,
    }
    h.SuccessJsonResponse(w, res)
    return
}
func (h *InvestmentHandlerChi) getTaxaSelic() string {
    vlr, err := h.bcbService.GetSelic()
    if err != nil {
        fmt.Println(err)
    }
    valueSelic := fmt.Sprintf("%.2f", vlr)
    valueSelic = strings.ReplaceAll(valueSelic, ".", ",")
    return valueSelic
}
func (h *InvestmentHandlerChi) getMediaIpca() string {
    vlr, err := h.bcbService.GetMediaIpca()
    if err != nil {
        fmt.Println(err)
    }
    valueIpca := fmt.Sprintf("%.2f", vlr)
    valueIpca = strings.ReplaceAll(valueIpca, ".", ",")
    return valueIpca
}
