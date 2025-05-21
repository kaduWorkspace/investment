package infra_investment

import (
	"kaduhod/fin_v3/core/domain/investment"
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	"time"
)

type FutureValueOfASerieDecimal struct {

}

func (self FutureValueOfASerieDecimal) Calculate(contribuition, taxDecimal, periods valueobjects.Money, firstDay bool) float64 {
    one := NewDecimalMoney(1.0)
    monthlyTax := self.monthlyTax(taxDecimal)
    growthFactor := one.Add(monthlyTax).Pow(periods).Subtract(one)
    growthFactorPerMonthlyTax := growthFactor.Divide(monthlyTax)
    if firstDay {
        growthFactorPerMonthlyTax = one.Add(monthlyTax).Multiply(growthFactorPerMonthlyTax)
    }
    result := contribuition.Multiply(growthFactorPerMonthlyTax)
    return result.GetAmount()
}
func (self FutureValueOfASerieDecimal) monthlyTax(tax valueobjects.Money) valueobjects.Money {
    twelve := NewDecimalMoney(12.0)
    return tax.Divide(twelve)
}
func (self FutureValueOfASerieDecimal) CalculateTrackingPeriods(initialValue, contribuition, taxDecimal, periods valueobjects.Money, firstDay bool, initialDate time.Time) (valueobjects.Money ,[]investment.PeriodTracker) {
    monthlyTax := self.monthlyTax(taxDecimal)
    accrued := initialValue
    counter := 0
    periodsTracker := []investment.PeriodTracker{}
    for counter < int(periods.GetAmount()) {
        if firstDay {
            accrued = accrued.Add(contribuition)
        }
        interest := accrued.Multiply(monthlyTax)
        accrued = interest.Add(accrued)
        if !firstDay {
            accrued = accrued.Add(contribuition)
        }
        periodsTracker = append(periodsTracker, investment.NewPeriodTracker(accrued, counter + 1, interest, initialDate))
        counter++
    }
    futureValue := accrued
    return futureValue, periodsTracker
}
