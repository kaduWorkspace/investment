package infra_investment

import (
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
			want:                  1268.25,
		},
		{
			name:                  "monthly contributions start period",
			interestRateDecimal:   0.12,
			periods:               12,
			contributionAmount:    100,
			contributionOnFirstDay: true,
			want:                  1280.93,
		},
		{
			name:                  "quarterly contributions low rate",
			interestRateDecimal:   0.01,
			periods:               4,
			contributionAmount:    500,
			contributionOnFirstDay: false,
			want:                  2002.50,
		},
		{
			name:                  "weekly contributions high rate",
			interestRateDecimal:   0.24,
			periods:               52,
			contributionAmount:    10,
			contributionOnFirstDay: true,
			want:                  918.16,
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
			if almostEqual(got, tt.want, 0.0000000000001) && got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
    testsWithInitialValue := []struct {
		name                  string
		interestRateDecimal   float64
		periods               int
		contributionAmount    float64
		contributionOnFirstDay bool
        initialValue          float64
		want                  float64
	}{
		{
			name:                  "loop_with initial value end period",
			interestRateDecimal:   0.12,
			periods:               12,
			contributionOnFirstDay: true,
			want:            2062.84,
			initialValue:          500.00,
		},
		{
			name:                  "loop_with initial value start period",
			interestRateDecimal:   0.12,
			periods:               12,
			contributionOnFirstDay: true,
			want:            2535.62,
			initialValue:          300.00,
		},
    }
    cp := CompoundInterestDecimal{}
    today := time.Now()
    for _, tt := range testsWithInitialValue {
		t.Run(tt.name, func(t *testing.T) {
			contribution := NewDecimalMoney(tt.contributionAmount)
			tax := NewDecimalMoney(tt.interestRateDecimal)
			periods := tt.periods
            initialValue := NewDecimalMoney(tt.initialValue)
            futureValue, _ := fv.CalculateTrackingPeriods(initialValue, contribution, tax, tt.contributionOnFirstDay, today, periods)
            compoundInterest := cp.Calculate(initialValue, tax, periods)
            result := futureValue.GetAmount() + compoundInterest
            if almostEqual(result, tt.want, 0.0000000000001) && result != tt.want {
				t.Errorf("Calculate() = %v, want %v", result, tt.want)
            }
		})
    }
}
