package app_investment_decimal

import (
    valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
    "testing"
    "time"
)

func TestFutureValueOfASerieDecimal_Calculate(t *testing.T) {
    tests := []struct {
        name                  string
        interestRateDecimal   float64
        periods               int
        contributionAmount    float64
        contributionOnFirstDay bool
        want                  float64
    }{
        {
            name:                  "monthly contributions end period",
            interestRateDecimal:   0.12,
            periods:               12,
            contributionAmount:    100,
            contributionOnFirstDay: false,
            want:                  1268.2503,
        },
        {
            name:                  "monthly contributions start period",
            interestRateDecimal:   0.12,
            periods:               12,
            contributionAmount:    100,
            contributionOnFirstDay: true,
            want:                  1280.9328,
        },
        {
            name:                  "quarterly contributions low rate",
            interestRateDecimal:   0.01,
            periods:               4,
            contributionAmount:    500,
            contributionOnFirstDay: false,
            want:                  2002.5013,
        },
        {
            name:                  "weekly contributions high rate",
            interestRateDecimal:   0.24,
            periods:               52,
            contributionAmount:    10,
            contributionOnFirstDay: true,
            want:                  918.1673,
        },
        {
            name:                  "single contribution edge case",
            interestRateDecimal:   0.05,
            periods:               1,
            contributionAmount:    1000,
            contributionOnFirstDay: false,
            want:                  1000.00,
        },
    }


    fv := FutureValueOfASerieDecimal{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            contribution := NewDecimalMoney(tt.contributionAmount)
            tax := NewDecimalMoney(tt.interestRateDecimal)
            periods := tt.periods

            got := fv.Calculate(contribution, tax, tt.contributionOnFirstDay, periods)
            if !almostEqual(got.GetAmount(), tt.want, 0.0001) {
                t.Errorf("Calculate() = %v, want %v", got, tt.want)
            }
        })
    }
    testsWithInitialValue := []struct {
        name                  string
        interestRateDecimal   float64
        interestRateDecimalInflation   float64
        periods               int
        contributionAmount    float64
        contributionOnFirstDay bool
        initialValue          float64
        want                  float64
        wantReal              float64
    }{
        {
            name:                  "loop_with initial value end period",
            interestRateDecimal:   0.12,
            interestRateDecimalInflation:   0.05,
            periods:               12,
            contributionOnFirstDay: true,
            contributionAmount:    117,
            want:            2062.1038,
            wantReal:        1990.1182,
            initialValue:          500.00,
        },
        {
            name:                  "loop_with initial value start period",
            interestRateDecimal:   0.12,
            interestRateDecimalInflation:   0.14,
            contributionAmount:    154,
            periods:               12,
            contributionOnFirstDay: true,
            want:            2310.6840,
            wantReal:        2125.3113,
            initialValue:          300.00,
        },
    }
    today := time.Now()
    for _, tt := range testsWithInitialValue {
        t.Run(tt.name, func(t *testing.T) {
            contribution := NewDecimalMoney(tt.contributionAmount)
            tax := NewDecimalMoney(tt.interestRateDecimal)
            taxInflation := NewDecimalMoney(tt.interestRateDecimalInflation)
            periods := tt.periods
            initialValue := NewDecimalMoney(tt.initialValue)
            futureValue, _ := fv.CalculateTrackingPeriods(initialValue, contribution, tax, tt.contributionOnFirstDay, today, periods)
            if !almostEqual(futureValue.GetAmount(), tt.want, 0.0001) {
                t.Errorf("Calculate() = %v, want %v", futureValue.GetAmount(), tt.want)
            }
            futureReal,_ := fv.CalculateTrackingPeriodsRealValue(initialValue, contribution, tax, taxInflation, tt.contributionOnFirstDay, today, periods)
            if !almostEqual(futureReal.GetAmount(), tt.wantReal, 0.0001) {
                t.Errorf("CalculateReal() = %v, want %v", futureReal.GetAmount(), tt.want)
            }
        })
    }
}
func TestFutureValueOfASerieDecimal_CalculateReal(t *testing.T) {
    tests := []struct {
        name                  string
        interestRateDecimal   float64
        inflationTax float64
        periods               int
        contributionAmount    float64
        contributionOnFirstDay bool
        want                  float64
        wantReal float64
    }{
        {
            name:                  "Value asjusted by inflation",
            interestRateDecimal:   0.12,
            inflationTax:   0.04,
            periods:               36,
            contributionAmount:    100,
            contributionOnFirstDay: false,
            want:                  4307.6878,
            wantReal:              4034.8028,
        },
        {
            name:                  "Value asjusted by inflation with bigger inflation",
            interestRateDecimal:   0.12,
            inflationTax:   0.14,
            periods:               36,
            contributionAmount:    100,
            contributionOnFirstDay: false,
            want:                  4307.6878,
            wantReal:              3509.4026,
        },
    }


    fv := FutureValueOfASerieDecimal{}
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            contribution := NewDecimalMoney(tt.contributionAmount)
            tax := NewDecimalMoney(tt.interestRateDecimal)
            inflationTax := NewDecimalMoney(tt.inflationTax)
            periods := tt.periods

            got := fv.Calculate(contribution, tax, tt.contributionOnFirstDay, periods)
            if !almostEqual(got.GetAmount(), tt.want, 0.0001) {
                t.Errorf("Calculate() = %v, want %v", got, tt.want)
            }
            gotReal := fv.CalculateRealValue(contribution, tax, inflationTax, tt.contributionOnFirstDay, periods)
            t.Log(gotReal.GetAmount(), got.GetAmount())
            if !almostEqual(gotReal.GetAmount(), tt.wantReal, 0.0001) {
                t.Errorf("CalculateReal() = %v, want %v", got, tt.want)
            }
        })
    }
}
func TestFutureValueOfASerieDecimal_PredictConstribuiton(t *testing.T) {
    tests := []struct {
        name                  string
        interestRateDecimal   float64
        interestRateDecimalInflation   float64
        periods               int
        finalValue            float64
        initialValue          float64
        contributionOnFirstDay bool
        want                  float64
        wantReal                  float64
    }{
        {
            name:                  "with initial value end period",
            interestRateDecimal:   0.12,
            interestRateDecimalInflation:   0.05,
            periods:               12,
            contributionOnFirstDay: true,
            finalValue:            2062.84,
            initialValue:          500.00,
            want:                  117.057,
            wantReal:                  122.8447,
        },
        {
            name:                  "with initial value start period",
            interestRateDecimal:   0.12,
            interestRateDecimalInflation:   0.05,
            periods:               12,
            contributionOnFirstDay: true,
            finalValue:            2535.62,
            initialValue:          300.00,
            want:                  171.56,
            wantReal:                  178.0217,
        },
        {
            name:                  "monthly contributions",
            interestRateDecimal:   0.12,
            interestRateDecimalInflation:   0.05,
            periods:               12,
            finalValue:            1280.93,
            want:                  100.99977888174818,
            wantReal:                  103.5216,
        },
        {
            name:                  "semester contributions",
            interestRateDecimal:   0.01,
            interestRateDecimalInflation:   0.05,
            periods:               6,
            finalValue:            2015.87,
            want:                  335.2790587002493,
            wantReal:                  338.6547,
        },
    }

    fv := FutureValueOfASerieDecimal{}
    today := time.Now()
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            finalValue := NewDecimalMoney(tt.finalValue)
            tax := NewDecimalMoney(tt.interestRateDecimal)
            var initialValue valueobjects.Money
            if tt.initialValue > 0 {
                initialValue = NewDecimalMoney(tt.initialValue)
            } else {
                initialValue = NewDecimalMoney(0)
            }
            inflationTax := NewDecimalMoney(tt.interestRateDecimalInflation)
            got := fv.PredictContribution(finalValue, tax, initialValue, tt.contributionOnFirstDay, tt.periods)
            if !almostEqual(got.GetAmount(), tt.want, 0.01) {
                t.Logf(`
                Test Case: %s
                Interest Rate: %.2f
                Periods: %d
                Final Value: %.2f
                Initial Value: %.2f
                Contribution On First Day: %v
                Want: %.3f`,
                tt.name,
                tt.interestRateDecimal,
                tt.periods,
                tt.finalValue,
                tt.initialValue,
                tt.contributionOnFirstDay,
                tt.want)
                t.Errorf("PredictConstribuiton() = %v, want %v", got.GetAmount(), tt.want)
            }
            gotReal := fv.PredictContributionRealValue(finalValue, tax, inflationTax, initialValue, tt.contributionOnFirstDay, tt.periods)
            if !almostEqual(gotReal.GetAmount(), tt.wantReal, 0.0001) {
                t.Errorf("PredictConstribuitonReal() = %v, want %v", gotReal.GetAmount(), tt.wantReal)
            }
            confirm,_ := fv.CalculateTrackingPeriods(initialValue ,got, tax, tt.contributionOnFirstDay, today, tt.periods)
            if !almostEqual(confirm.GetAmount(), tt.finalValue, 0.01) {
                t.Logf(`
                Test Case: %s
                Interest Rate: %.2f
                Periods: %d
                Final Value: %.2f
                Initial Value: %.2f
                Contribution On First Day: %v
                Want: %.3f
                contribution: %.3f`,
                tt.name,
                tt.interestRateDecimal,
                tt.periods,
                tt.finalValue,
                tt.initialValue,
                tt.contributionOnFirstDay,
                tt.want, got.GetAmount())
                t.Errorf("PredictConstribuiton() = %v, want %v", confirm.GetAmount(), tt.finalValue)
            }
        })
    }
}
