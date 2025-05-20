package investment

import (
	"time"
	"fmt"
)
type Investment struct {
    interestRateDecimal float64
    periods             float64
    contributionAmount  float64
}

// GetInterestRateDecimal returns the annual interest rate in decimal form
func (i *Investment) GetInterestRateDecimal() float64 {
    return i.interestRateDecimal
}

// SetInterestRateDecimal sets the annual interest rate in decimal form
func (i *Investment) SetInterestRateDecimal(rate float64) {
    i.interestRateDecimal = rate
}

// GetPeriods returns the number of periods
func (i *Investment) GetPeriods() float64 {
    return i.periods
}

// SetPeriods sets the number of periods
func (i *Investment) SetPeriods(periods float64) {
    i.periods = periods
}

// GetContributionAmount returns the contribution amount
func (i *Investment) GetContributionAmount() float64 {
    return i.contributionAmount
}

// SetContributionAmount sets the contribution amount
func (i *Investment) SetContributionAmount(amount float64) {
    i.contributionAmount = amount
}
func (i Investment) GetDates(initialDate, finalDate string) ([]time.Time, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, initialDate)
	if err != nil {
		return nil, fmt.Errorf("invalid initial date: %v", err)
	}
	end, err := time.Parse(layout, finalDate)
	if err != nil {
		return nil, fmt.Errorf("invalid final date: %v", err)
	}

	if end.Before(start) {
		return nil, fmt.Errorf("final date must be after initial date")
	}
    months := []time.Time{}
	for start.Before(end) {
		start = start.AddDate(0, 1, 0)
        months = append(months, start)
	}

	return months, nil
}
func (i Investment) MonthsBetweenDates(initialDate, finalDate string) (int, error) {
	layout := "2006-01-02"
	start, err := time.Parse(layout, initialDate)
	if err != nil {
		return 0, fmt.Errorf("invalid initial date: %v", err)
	}
	end, err := time.Parse(layout, finalDate)
	if err != nil {
		return 0, fmt.Errorf("invalid final date: %v", err)
	}

	if end.Before(start) {
		return 0, fmt.Errorf("final date must be after initial date")
	}

	months := 0
	for start.Before(end) {
		start = start.AddDate(0, 1, 0)
		months++
	}

	return months, nil
}
