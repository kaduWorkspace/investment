package interface_chi

import (
	"encoding/json"
	"errors"
	"fmt"
	app_investment_decimal "kaduhod/fin_v3/core/application/investment/service/decimal"
	"kaduhod/fin_v3/core/domain/external"
	core_http "kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/investment"
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	validators_dto "kaduhod/fin_v3/core/interfaces/http/dto/validators"
	"kaduhod/fin_v3/core/interfaces/web/renderer"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"strconv"
	"strings"
	"time"
)
type InvestmentHandlerChiWeb struct {
    bcbService external.BcbI
    CompoundInterestService investment.CompoundInterest
    FutureValueOfASeriesService investment.FutureValueOfASeries
    Renderer *renderer.Renderer
    sessionService core_http.SessionService
}
func NewInvestmentHandlerChiWeb(bcb external.BcbI ,sessionService core_http.SessionService ,compoundInterestService investment.CompoundInterest, futureValueOfASeriesService investment.FutureValueOfASeries, renderer *renderer.Renderer) core_http.InvestmentHandlerWeb {
    return &InvestmentHandlerChiWeb{
        CompoundInterestService: compoundInterestService,
        FutureValueOfASeriesService: futureValueOfASeriesService,
        Renderer: renderer,
        sessionService: sessionService,
        bcbService: bcb,
    }
}
func (h *InvestmentHandlerChiWeb) Index(w http.ResponseWriter, r *http.Request) {
    if err := h.Renderer.Render(w, "base", nil); err != nil {
        w.WriteHeader(500)
    }
}
func (h *InvestmentHandlerChiWeb) GetSessionService() core_http.SessionService {
    return h.sessionService
}
func (h *InvestmentHandlerChiWeb) getSession(r *http.Request) (core_http.SessionData, error) {
    session, err := h.sessionService.Get(struct_utils.GetCookie(r).Value)
    if err != nil {
        fmt.Println(err)
        return session, err
    }
    return session, nil
}
func (h *InvestmentHandlerChiWeb) getCsrfToken(r *http.Request) (string, error) {
    session, err := h.getSession(r)
    if err != nil {
        return "", errors.New("Session not found")
    }
    return session.Csrf, nil
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesPredictFormPage(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    data := map[string]any{
        "csrf": csrf,
        "selic_tax": h.getTaxaSelic(),
        "ipca_media": h.getMediaIpca(),
    }
    if err := h.Renderer.Render(w, "fv_predict_form_result_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesPredictResultPage(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    initialValueF, err := strconv.ParseFloat(r.FormValue("initial_value"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid initial_value"))
        return
    }
    firstDay, err := strconv.ParseBool(r.FormValue("first_day"))
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid first_day"))
        return
    }
    taxF, err := strconv.ParseFloat(r.FormValue("tax_decimal"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid tax_decimal"))
        return
    }
    taxF = taxF / 100

    taxInflationF, err := strconv.ParseFloat(r.FormValue("tax_decimal_inflation"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid tax_decimal_inflation"))
        return
    }
    taxInflationF = taxInflationF / 100
    periodsF, err := strconv.ParseInt(r.FormValue("periods"), 10, 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid periods"))
        return
    }
    finalValueF, err := strconv.ParseFloat(r.FormValue("final_value"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid final value"))
        return
    }
    userInput := validators_dto.PredictContributionFVSInput{
        InitialValue: initialValueF,
        ContributionOnFirstDay: firstDay,
        Periods: int(periodsF),
        TaxDecimal: taxF,
        FinalValue: finalValueF,
        TaxDecimalInflation: taxInflationF,
    }
    if err := userInput.Validate(userInput); err != nil {
        fmt.Println(err)
        errs := userInput.FormatValidationError(err, "pt")
        w.Header().Set("HX-Retarget", "#form_container")
        h.Renderer.Render(w, "fv_predict_form", map[string]any{
            "csrf": csrf,
            "selic_tax": h.getTaxaSelic(),
            "ipca_media": h.getMediaIpca(),
            "errs": errs,
        })
        return
    }
    finalValue := app_investment_decimal.NewDecimalMoney(userInput.FinalValue)
    taxDecimal := app_investment_decimal.NewDecimalMoney(userInput.TaxDecimal)
    initialValue := app_investment_decimal.NewDecimalMoney(userInput.InitialValue)
    contribution := h.FutureValueOfASeriesService.PredictContribution(
        finalValue,
        taxDecimal,
        initialValue,
        userInput.ContributionOnFirstDay,
        userInput.Periods,
    )
    taxInflation := app_investment_decimal.NewDecimalMoney(userInput.TaxDecimalInflation)
    contributionReal := h.FutureValueOfASeriesService.PredictContributionRealValue(
        finalValue,
        taxDecimal,
        taxInflation,
        initialValue,
        userInput.ContributionOnFirstDay,
        userInput.Periods,
    )
    one := app_investment_decimal.NewDecimalMoney(1.0)
    hundred := app_investment_decimal.NewDecimalMoney(100.0)
    taxReal := one.Add(taxDecimal).Divide(one.Add(taxInflation)).Subtract(one).Multiply(hundred).Formatted()
    data := map[string]any{
        "csrf": csrf,
        "selic_tax": h.getTaxaSelic(),
        "ipca_media": h.getMediaIpca(),
        "final_value": finalValue.Formatted(),
        "contribution_needed": contribution.Formatted(),
        "contribution_needed_real": contributionReal.Formatted(),
        "initial_value": initialValue.Formatted(),
        "tax": taxDecimal.Multiply(hundred).Formatted(),
        "tax_real": taxReal,
    }
    if err := h.Renderer.Render(w, "predict_result", data); err != nil {
        fmt.Println(err)
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesFormPage(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    data := map[string]any{
        "csrf": csrf,
        "selic_tax": h.getTaxaSelic(),
        "ipca_media": h.getMediaIpca(),
    }
    if err := h.Renderer.Render(w, "fv_form_result_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesResultPage(w http.ResponseWriter, r *http.Request) {
    csrf, err := h.getCsrfToken(r)
    if err != nil {
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    initialValueF, err := strconv.ParseFloat(r.FormValue("initial_value"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid initial_value"))
        return
    }
    firstDay, err := strconv.ParseBool(r.FormValue("first_day"))
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid fist_day"))
        return
    }
    contributionF, err := strconv.ParseFloat(r.FormValue("contribution"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid contribution"))
        return
    }
    taxF, err := strconv.ParseFloat(r.FormValue("tax_decimal"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid tax_decimal"))
        return
    }
    taxF = taxF / 100
    periodsF, err := strconv.ParseInt(r.FormValue("periods"), 10, 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid periods"))
        return
    }
    taxInflationF, err := strconv.ParseFloat(r.FormValue("tax_decimal_inflation"), 64)
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid tax_decimal_inflation"))
        return
    }
    taxInflationF = taxInflationF / 100
    if err != nil {
        w.WriteHeader(400)
        w.Write([]byte("invalid periods"))
        return
    }
    userInput := validators_dto.FutureValueOfASeriesWithPeriodsInput{
        InitialValue: initialValueF,
        FirstDay: firstDay,
        Periods: int(periodsF),
        TaxDecimal: taxF,
        Contribution: contributionF,
        InitialDate: r.FormValue("initial_date"),
        TaxDecimalInflation: taxInflationF,
    }
    if err := userInput.Validate(userInput); err != nil {
        fmt.Println(err)
        errs := userInput.FormatValidationError(err, "pt")
        w.Header().Set("HX-Retarget", "#form_container")
        h.Renderer.Render(w, "fv_form", map[string]any{
            "csrf": csrf,
            "selic_tax": h.getTaxaSelic(),
            "ipca_media": h.getMediaIpca(),
            "errs": errs,
        })
        return
    }
    initialValue := app_investment_decimal.NewDecimalMoney(userInput.InitialValue)
    contribution := app_investment_decimal.NewDecimalMoney(userInput.Contribution)
    periodsD := app_investment_decimal.NewDecimalMoney(float64(userInput.Periods))
    taxDecimal := app_investment_decimal.NewDecimalMoney(userInput.TaxDecimal)
    result, periods := h.FutureValueOfASeriesService.CalculateTrackingPeriods(
        initialValue,
        contribution,
        taxDecimal,
        userInput.FirstDay,
        time.Now(),
        userInput.Periods,
    )
    periods = setupItensFromPeriods(periods, struct_utils.EhMobile(r.UserAgent()))
    b, err := json.Marshal(periods)
    var table string
    if err != nil {
        fmt.Println("Error building json table of period trackers")
    } else {
        table = string(b)
    }
    taxInflation := app_investment_decimal.NewDecimalMoney(userInput.TaxDecimalInflation)
    resultReal, periodsReal := h.FutureValueOfASeriesService.CalculateTrackingPeriodsRealValue(
        initialValue,
        contribution,
        taxDecimal,
        taxInflation,
        userInput.FirstDay,
        time.Now(),
        userInput.Periods,
    )
    periodsReal = setupItensFromPeriods(periodsReal, struct_utils.EhMobile(r.UserAgent()))
    b, err = json.Marshal(periodsReal)
    var tableReal string
    if err != nil {
        fmt.Println("Error building json table of period trackers")
    } else {
        tableReal = string(b)
    }
    totalInvested := periodsD.Multiply(contribution).Add(initialValue)
    var initialValueOrOne valueobjects.Money
    if userInput.InitialValue < 1 {
        initialValueOrOne = app_investment_decimal.NewDecimalMoney(1.0)
    } else {
        initialValueOrOne = initialValue
    }
    one := app_investment_decimal.NewDecimalMoney(1.0)
    hundred := app_investment_decimal.NewDecimalMoney(100.0)
    taxReal := one.Add(taxDecimal).Divide(one.Add(taxInflation)).Subtract(one).Multiply(hundred).Formatted()
    roi := result.Subtract(app_investment_decimal.NewDecimalMoney(userInput.InitialValue))
    roiReal := resultReal.Subtract(app_investment_decimal.NewDecimalMoney(userInput.InitialValue))
    roiPorcentage := roi.Divide(initialValueOrOne).Multiply(app_investment_decimal.NewDecimalMoney(100))
    roiPorcentageReal := roiReal.Divide(initialValueOrOne).Multiply(app_investment_decimal.NewDecimalMoney(100))
    netGain := result.Subtract(periodsD.Multiply(contribution))
    netGainReal := resultReal.Subtract(periodsD.Multiply(contribution))
    data := map[string]any{
        "csrf": csrf,
        "selic_tax": h.getTaxaSelic(),
        "ipca_media": h.getMediaIpca(),
        "periods_json": table,
        "periods_real_json": tableReal,
        "roi": roi.Formatted(),// return of investment | valorizacao
        "roi_real": roiReal.Formatted(),// return of investment | valorizacao
        "total_invested": totalInvested.Formatted(),// total investido
        "initial_value": initialValue.Formatted(),
        "final_value": result.Formatted(),
        "final_value_real": resultReal.Formatted(),
        "net_gain": netGain.Formatted(),// juros rendido | rentabilidade liquida.
        "net_gain_real": netGainReal.Formatted(),// juros rendido | rentabilidade liquida.
        "roi_porcentage": roiPorcentage.Formatted(), // retorno sobre o investimento
        "roi_porcentage_real": roiPorcentageReal.Formatted(), // retorno sobre o investimento
        "contribution": contribution.Formatted(),
        "periodsTracker": periods,
        "periodsTrackerReal": periodsReal,
        "tax_real": taxReal,
        "tax": taxDecimal.Multiply(hundred).Formatted(),
    }
    if err := h.Renderer.Render(w, "fv_result", data); err != nil {
        fmt.Println(err)
    }
}
func setupItensFromPeriods(periods []investment.PeriodTracker, for_mobile bool) []investment.PeriodTracker {
    var max_table_items int
    if for_mobile {
        max_table_items = 12
    } else {
        max_table_items = 20
    }
    data_count := len(periods)
    if data_count <= max_table_items {
        return periods
    }
    adjusted_table := make([]investment.PeriodTracker, 0, max_table_items)
    step := int(data_count / max_table_items)
    count := 0
    var curr investment.PeriodTracker
    for len(adjusted_table) < max_table_items {
        if len(adjusted_table) == 0 {
            curr = periods[0]
        } else if len(adjusted_table) == max_table_items - 1 {
            curr = periods[len(periods) - 1]
        } else {
            curr = periods[count + step]
        }
        count = count + step
        adjusted_table = append(adjusted_table, curr)
    }
    return adjusted_table
}
func (h *InvestmentHandlerChiWeb) getTaxaSelic() string {
    vlr, err := h.bcbService.GetSelic()
    if err != nil {
        fmt.Println(err)
    }
    valueSelic := fmt.Sprintf("%.2f", vlr)
    valueSelic = strings.ReplaceAll(valueSelic, ".", ",")
    return valueSelic
}
func (h *InvestmentHandlerChiWeb) getMediaIpca() string {
    vlr, err := h.bcbService.GetMediaIpca()
    if err != nil {
        fmt.Println(err)
    }
    valueIpca := fmt.Sprintf("%.2f", vlr)
    valueIpca = strings.ReplaceAll(valueIpca, ".", ",")
    return valueIpca
}
