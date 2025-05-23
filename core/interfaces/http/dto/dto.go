package http_dto
import "time"

type InvestmentInput struct {
    Periods int `json:"periods" form:"periods"`
    TaxDecimal float64 `json:"tax_decimal" form:"tax_decimal"`
}

type SeriesInvestmentInput struct {
    FirstDay bool `json:"first_day" form:"first_day"`
    Contribution float64 `json:"contribution" form:"contribution"`
}

type CoumpoundInterestInput struct {
    InvestmentInput
    InitialValue float64 `json:"initial_value" form:"initial_value"`
}

type FutureValueOfASeriesInput struct {
    InvestmentInput
    SeriesInvestmentInput
}

type FutureValueOfASeriesWithPeriodsInput struct {
    InvestmentInput
    SeriesInvestmentInput
    InitialValue float64 `json:"initial_value" form:"initial_value"`
    InitialDate time.Time `json:"initial_date" form:"initial_date"`
}

type PredictContributionFVSInput struct {
    InvestmentInput
    FinalValue float64 `json:"final_value" form:"final_value"`
    InitialValue float64 `json:"initial_value" form:"initial_value"`
    ContributionOnFirstDay bool `json:"contribution_on_first_day" form:"contribution_on_first_day"`
}

