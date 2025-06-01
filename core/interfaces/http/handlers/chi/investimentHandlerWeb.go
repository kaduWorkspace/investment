package interface_chi

import (
	"encoding/json"
	"fmt"
	core_http "kaduhod/fin_v3/core/domain/http"
	"kaduhod/fin_v3/core/domain/investment"
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	infra_investment "kaduhod/fin_v3/core/infra/investment/decimal"
	validators_dto "kaduhod/fin_v3/core/interfaces/http/dto/validators"
	"kaduhod/fin_v3/core/interfaces/web/renderer"
	struct_utils "kaduhod/fin_v3/pkg/utils/struct"
	"net/http"
	"strconv"
	"strings"
	"time"
)
type InvestmentHandlerChiWeb struct {
    CompoundInterestService investment.CompoundInterest
    FutureValueOfASeriesService investment.FutureValueOfASeries
    Renderer *renderer.Renderer
}
func NewInvestmentHandlerChiWeb(compoundInterestService investment.CompoundInterest, futureValueOfASeriesService investment.FutureValueOfASeries, renderer *renderer.Renderer) core_http.InvestmentHandlerWeb {
    return &InvestmentHandlerChiWeb{
        CompoundInterestService: compoundInterestService,
        FutureValueOfASeriesService: futureValueOfASeriesService,
        Renderer: renderer,
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesPredictFormPage(w http.ResponseWriter, r *http.Request) {
    data := map[string]any{
        "csrf": "1234546",
        "selic_tax": h.GetTaxaSelic(),
    }
    if err := h.Renderer.Render(w, "fv_predict_form_result_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesPredictResultPage(w http.ResponseWriter, r *http.Request) {
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
    }
    if err := userInput.Validate(userInput); err != nil {
        fmt.Println(err)
        errs := userInput.FormatValidationError(err, "pt")
        w.Header().Set("HX-Retarget", "#form_container")
        h.Renderer.Render(w, "fv_predict_form", map[string]any{
            "csrf": "1234546",
            "selic_tax": h.GetTaxaSelic(),
            "errs": errs,
        })
        return
    }
    finalValue := infra_investment.NewDecimalMoney(userInput.FinalValue)
    taxDecimal := infra_investment.NewDecimalMoney(userInput.TaxDecimal)
    initialValue := infra_investment.NewDecimalMoney(userInput.InitialValue)
    contribution := h.FutureValueOfASeriesService.PredictContribution(
        finalValue,
        taxDecimal,
        initialValue,
        userInput.ContributionOnFirstDay,
        userInput.Periods,
    )
    data := map[string]any{
        "csrf": "1234546",
        "selic_tax": h.GetTaxaSelic(),
        "final_value": finalValue.Formatted(),
        "contribution_needed": contribution.Formatted(),
        "initial_value": initialValue.Formatted(),
    }
    if err := h.Renderer.Render(w, "predict_result", data); err != nil {
        fmt.Println(err)
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesFormPage(w http.ResponseWriter, r *http.Request) {
    data := map[string]any{
        "csrf": "1234546",
        "selic_tax": h.GetTaxaSelic(),
    }
    if err := h.Renderer.Render(w, "fv_form_result_page", data); err != nil {
        fmt.Println(err)
    }
}
func (h *InvestmentHandlerChiWeb) FutureValueOfASeriesResultPage(w http.ResponseWriter, r *http.Request) {
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
    }
    if err := userInput.Validate(userInput); err != nil {
        fmt.Println(err)
        errs := userInput.FormatValidationError(err, "pt")
        w.Header().Set("HX-Retarget", "#form_container")
        h.Renderer.Render(w, "fv_form", map[string]any{
            "csrf": "1234546",
            "selic_tax": h.GetTaxaSelic(),
            "errs": errs,
        })
        return
    }
    initialValue := infra_investment.NewDecimalMoney(userInput.InitialValue)
    contribution := infra_investment.NewDecimalMoney(userInput.Contribution)
    periodsD := infra_investment.NewDecimalMoney(float64(userInput.Periods))
    result, periods := h.FutureValueOfASeriesService.CalculateTrackingPeriods(
        initialValue,
        contribution,
        infra_investment.NewDecimalMoney(userInput.TaxDecimal),
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
    totalInvested := periodsD.Multiply(contribution).Add(initialValue)
    var initialValueOrOne valueobjects.Money
    if userInput.InitialValue < 1 {
        initialValueOrOne = infra_investment.NewDecimalMoney(1.0)
    } else {
        initialValueOrOne = initialValue
    }
    roi := result.Subtract(infra_investment.NewDecimalMoney(userInput.InitialValue))
    roiPorcentage := roi.Divide(initialValueOrOne).Multiply(infra_investment.NewDecimalMoney(100))
    netGain := result.Subtract(periodsD.Multiply(contribution))
    data := map[string]any{
        "csrf": "1234546",
        "selic_tax": h.GetTaxaSelic(),
        "periods_json": table,
        "roi": roi.Formatted(),// return of investment | valorizacao
        "total_invested": totalInvested.Formatted(),// total investido
        "initial_value": initialValue.Formatted(),
        "final_value": result.Formatted(),
        "net_gain": netGain.Formatted(),// juros rendido | rentabilidade liquida.
        "roi_porcentage": roiPorcentage.Formatted(), // retorno sobre o investimento
        "contribution": contribution.Formatted(),
        "periodsTracker": periods,
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
func (h *InvestmentHandlerChiWeb) GetTaxaSelic() string {
    valueSelic := "13,25" // default
    result, err := struct_utils.HttpRequest("https://www.bcb.gov.br/api/servico/sitebcb//taxaselic/ultima?withCredentials=true", "GET",
        map[string]string{"content-type":"text/plain"}, "")
    if err != nil {
        return valueSelic
    }
    var response map[string]interface{}
    err, response = struct_utils.FromJson[map[string]interface{}]([]byte(result))
    if err != nil {
        return valueSelic
    }
    content, ok := response["conteudo"].([]interface{})
    if !ok || len(content) == 0 {
        return valueSelic
    }
    firstItem, ok := content[0].(map[string]interface{})
    if !ok {
        return valueSelic
    }
    if metaSelic, ok := firstItem["MetaSelic"].(float64); ok {
        valueSelic = fmt.Sprintf("%.2f", metaSelic)
        valueSelic = strings.ReplaceAll(valueSelic, ".", ",")
    }
    return valueSelic
}
