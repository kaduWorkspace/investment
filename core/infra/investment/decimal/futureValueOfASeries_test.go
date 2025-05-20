package infra_investment

import (
	"testing"
)

func TestFutureValueOfASerieDecimal_Calculate(t *testing.T) {
	tests := []struct {
		name                  string
		interestRateDecimal   float64
		periods               float64
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
			periods := NewDecimalMoney(tt.periods)

			got := fv.Calculate(contribution, tax, periods, tt.contributionOnFirstDay)
			if almostEqual(got, tt.want, 0.0000000000001) && got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
