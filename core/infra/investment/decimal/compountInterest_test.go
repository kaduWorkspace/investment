package infra_investment

import (
	"math"
	"testing"
)

func TestCompoundInterestDecimal_Calculate(t *testing.T) {
	tests := []struct {
		name         string
		initialValue float64
		tax          float64
		periods      int
		want         float64
	}{

        {
            name:     "Small investment with low interest",
            initialValue:  1000.0,
            tax:      0.01,
            periods:   12,
            want: 1010.045,
        },
        {
            name:     "Large investment with moderate interest",
            initialValue:  10000.0,
            tax:      0.05,
            periods:   24,
            want: 11049.41,
        },
        {
            name:     "Small investment with high interest",
            initialValue:  500.0,
            tax:      0.2,
            periods:   6,
            want: 552.13,
        },
        {
            name:     "Large investment with small monthly gain",
            initialValue:  50000.0,
            tax:      0.001,
            periods:   36,
            want: 50150.21,
        },
        {
            name:     "Zero months should return initial value",
            initialValue:  1000.0,
            tax:      0.1,
            periods:   0,
            want: 1000.0,
        },
	}
    cp := CompoundInterestDecimal{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initial := NewDecimalMoney(tt.initialValue)
			tax := NewDecimalMoney(tt.tax)
			periods := tt.periods

			got := cp.Calculate(initial, tax, periods)
			if !almostEqual(got.GetAmount(), tt.want, 0.01) {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
func almostEqual(a, b, tolerance float64) bool {
    return math.Abs(a-b) <= tolerance || math.Abs(b-a) <= tolerance
}
