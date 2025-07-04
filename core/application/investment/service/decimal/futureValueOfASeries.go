package app_investment_decimal

import (
	"kaduhod/fin_v3/core/domain/investment"
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	"time"
)

type FutureValueOfASerieDecimal struct {}

func (self FutureValueOfASerieDecimal) Calculate(contribution, taxDecimal valueobjects.Money, firstDay bool, periods int) valueobjects.Money {
    one := NewDecimalMoney(1.0)
    monthlyTax := self.monthlyTax(taxDecimal)
    growthFactor := one.Add(monthlyTax).Pow(NewDecimalMoney(float64(periods))).Subtract(one)
    growthFactorPerMonthlyTax := growthFactor.Divide(monthlyTax)
    if firstDay {
        growthFactorPerMonthlyTax = one.Add(monthlyTax).Multiply(growthFactorPerMonthlyTax)
    }
    result := contribution.Multiply(growthFactorPerMonthlyTax)
    return result
}
func (self FutureValueOfASerieDecimal) CalculateRealValue(contribution, taxDecimal, inflationTax valueobjects.Money, firstDay bool, periods int) valueobjects.Money {
    return self.Calculate(contribution, self.taxAdjusted(taxDecimal, inflationTax), firstDay, periods)
}
func (self FutureValueOfASerieDecimal) taxAdjusted(taxDecimal, inflationTax valueobjects.Money) valueobjects.Money {
    one := NewDecimalMoney(1.0)
    return one.Add(taxDecimal).Divide(one.Add(inflationTax)).Subtract(one)
}
func (self FutureValueOfASerieDecimal) monthlyTax(tax valueobjects.Money) valueobjects.Money {
    twelve := NewDecimalMoney(12.0)
    return tax.Divide(twelve)
}
func (self FutureValueOfASerieDecimal) CalculateTrackingPeriods(initialValue, contribution, taxDecimal valueobjects.Money, firstDay bool, initialDate time.Time, periods int) (valueobjects.Money ,[]investment.PeriodTracker) {
    monthlyTax := self.monthlyTax(taxDecimal)
    accrued := initialValue
    counter := 0
    periodsTracker := []investment.PeriodTracker{}
    for counter < periods {
        if firstDay {
            accrued = accrued.Add(contribution)
        }
        interest := accrued.Multiply(monthlyTax)
        accrued = interest.Add(accrued)
        if !firstDay {
            accrued = accrued.Add(contribution)
        }
        periodsTracker = append(periodsTracker, investment.NewPeriodTracker(accrued, counter + 1, interest, initialDate))
        counter++
    }
    futureValue := accrued
    return futureValue, periodsTracker
}
func (self FutureValueOfASerieDecimal) CalculateTrackingPeriodsRealValue(initialValue, contribution, taxDecimal, taxInflation valueobjects.Money, firstDay bool, initialDate time.Time, periods int) (valueobjects.Money ,[]investment.PeriodTracker) {
    return self.CalculateTrackingPeriods(initialValue, contribution, self.taxAdjusted(taxDecimal, taxInflation), firstDay, initialDate, periods)
}
func (self FutureValueOfASerieDecimal) PredictContribution(finalValue, taxDecimal, initialValue valueobjects.Money, contributionOnFirstDay bool, periods int) (valueobjects.Money) {
    taxByMonths := self.monthlyTax(taxDecimal)
    one := NewDecimalMoney(1.0)
    growthFactor := one.Add(taxByMonths).Pow(NewDecimalMoney(float64(periods))).Subtract(one).Divide(taxByMonths)
    if contributionOnFirstDay {
        growthFactor = growthFactor.Multiply(one.Add(taxByMonths))
    }
    if initialValue.GetAmount() > 0.0 {
        cp := CompoundInterestDecimal{}
        finalValue = finalValue.Subtract(cp.Calculate(initialValue, taxDecimal, periods))
    }
    result := finalValue.Divide(growthFactor)
    return result
}
func (self FutureValueOfASerieDecimal) PredictContributionRealValue(finalValue, taxDecimal, inflationTax, initialValue valueobjects.Money, contributionOnFirstDay bool, periods int) (valueobjects.Money) {
    return self.PredictContribution(finalValue, self.taxAdjusted(taxDecimal, inflationTax), initialValue, contributionOnFirstDay, periods)
}
func NewFutureValueOfASerieDecimal() investment.FutureValueOfASeries {
    return &FutureValueOfASerieDecimal{}
}
