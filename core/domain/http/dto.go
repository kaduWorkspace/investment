package http

import "time"

type InvestmentInput struct {
    Periods int
    TaxDecimal float64
}
type SeriesInvestmentInput struct {
    FirstDay bool
    Contribution float64
}
type CoumpoundInterestInput struct {
    InvestmentInput
    InitialValue float64
}
type FutureValueOfASeriesInput struct {
    InvestmentInput
    SeriesInvestmentInput
}
type FutureValueOfASeriesWithPeriodsInput struct {
    InvestmentInput
    SeriesInvestmentInput
    InitialValue float64
    InitialDate time.Time
}
type PredictContributionFVSInputs struct {
    InvestmentInput
    FinalValue float64
    InitialValue float64
    ContributionOnFirstDay bool
}

