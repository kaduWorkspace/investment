package infra_investment

import (
	"testing"
	"github.com/shopspring/decimal"
)

func TestNewDecimalMoney(t *testing.T) {
	tests := []struct {
		name   string
		amount float64
	}{
		{"positive value", 123.45},
		{"negative value", -67.89},
		{"zero value", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewDecimalMoney(tt.amount)
			if got.GetAmount() != tt.amount {
				t.Errorf("NewDecimalMoney() = %v, want %v", got.GetAmount(), tt.amount)
			}
		})
	}
}

func TestDecimalMoneyOperations(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		add      float64
		sub      float64
		mul      float64
		div      float64
		pow      float64
	}{
		{
			name:     "basic operations",
			a:        10.0,
			b:        15.0,
			add:      25.0,
			sub:      -5.0,
			mul:      150.0,
			div:      0.6666666666666667,
			pow:      1000000000000000.0,
		},
		{
			name:     "decimal precision",
			a:        2.0,
			b:        3.0,
			add:      5.0,
			sub:      -1.0,
			mul:      6.0,
			div:      0.6666666666666667,
			pow:      8.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moneyA := NewDecimalMoney(tt.a)
			moneyB := NewDecimalMoney(tt.b)

			// Test Add
			if got := moneyA.Add(moneyB).GetAmount(); !decimal.NewFromFloat(got).Equal(decimal.NewFromFloat(tt.add)) {
				t.Errorf("Add() = %v, want %v", got, tt.add)
			}

			// Test Subtract
			if got := moneyA.Subtract(moneyB).GetAmount(); !decimal.NewFromFloat(got).Equal(decimal.NewFromFloat(tt.sub)) {
				t.Errorf("Subtract() = %v, want %v", got, tt.sub)
			}

			// Test Multiply
			if got := moneyA.Multiply(moneyB).GetAmount(); !decimal.NewFromFloat(got).Equal(decimal.NewFromFloat(tt.mul)) {
				t.Errorf("Multiply() = %v, want %v", got, tt.mul)
			}

			// Test Divide
			if got := moneyA.Divide(moneyB).GetAmount(); !decimal.NewFromFloat(got).Equal(decimal.NewFromFloat(tt.div)) {
				t.Errorf("Divide() = %v, want %v", got, tt.div)
			}

			// Test Pow
			if got := moneyA.Pow(moneyB).GetAmount(); !decimal.NewFromFloat(got).Equal(decimal.NewFromFloat(tt.pow)) {
				t.Errorf("Pow() = %v, want %v", got, tt.pow)
			}
		})
	}
}

func TestDecimalMoneyEquals(t *testing.T) {
	tests := []struct {
		name   string
		a      float64
		b      float64
		expect bool
	}{
		{"equal values", 123.45, 123.45, true},
		{"different values", 123.45, 543.21, false},
		{"zero values", 0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moneyA := NewDecimalMoney(tt.a)
			moneyB := NewDecimalMoney(tt.b)
			if got := moneyA.Equals(moneyB); got != tt.expect {
				t.Errorf("Equals() = %v, want %v", got, tt.expect)
			}
		})
	}
}
