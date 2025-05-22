package infra_investment

import (
	"kaduhod/fin_v3/core/domain/investment"
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	"time"
)

type FutureValueOfASerieDecimal struct {

}

func (self FutureValueOfASerieDecimal) Calculate(contribuition, taxDecimal valueobjects.Money, firstDay bool, periods int) valueobjects.Money {
    one := NewDecimalMoney(1.0)
    monthlyTax := self.monthlyTax(taxDecimal)
    growthFactor := one.Add(monthlyTax).Pow(NewDecimalMoney(float64(periods))).Subtract(one)
    growthFactorPerMonthlyTax := growthFactor.Divide(monthlyTax)
    if firstDay {
        growthFactorPerMonthlyTax = one.Add(monthlyTax).Multiply(growthFactorPerMonthlyTax)
    }
    result := contribuition.Multiply(growthFactorPerMonthlyTax)
    return result
}
func (self FutureValueOfASerieDecimal) monthlyTax(tax valueobjects.Money) valueobjects.Money {
    twelve := NewDecimalMoney(12.0)
    return tax.Divide(twelve)
}
func (self FutureValueOfASerieDecimal) CalculateTrackingPeriods(initialValue, contribuition, taxDecimal valueobjects.Money, firstDay bool, initialDate time.Time, periods int) (valueobjects.Money ,[]investment.PeriodTracker) {
    monthlyTax := self.monthlyTax(taxDecimal)
    accrued := initialValue
    counter := 0
    periodsTracker := []investment.PeriodTracker{}
    for counter < periods {
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
func (self FutureValueOfASerieDecimal) PredictConstribuiton(finalValue, taxDecimal, initialValue valueobjects.Money, contribuitionOnFirstDay bool, periods int) (valueobjects.Money) {
    //fmt.Println("Parameters: finalValue:", finalValue, "taxDecimal:", taxDecimal, "initialValue:", initialValue, "contribuitionOnFirstDay:", contribuitionOnFirstDay, "periods:", periods)
    taxByMonths := self.monthlyTax(taxDecimal)
    //fmt.Println("taxByMonths:", taxByMonths)

    one := NewDecimalMoney(1.0)
    //fmt.Println("one:", one)

    growthFactor := one.Add(taxByMonths).Pow(NewDecimalMoney(float64(periods))).Subtract(one).Divide(taxByMonths)
    //fmt.Println("initial growthFactor:", growthFactor)

    if contribuitionOnFirstDay {
        growthFactor = growthFactor.Multiply(one.Add(taxByMonths))
        //fmt.Println("adjusted growthFactor (contribuitionOnFirstDay):", growthFactor)
    }
    if initialValue.GetAmount() > 0.0 {
        cp := CompoundInterestDecimal{}
        finalValue = finalValue.Subtract(cp.Calculate(initialValue, taxDecimal, periods))
        //fmt.Println("finalValue after initialValue adjustment:", finalValue)
    }

    result := finalValue.Divide(growthFactor)
    //fmt.Println("final result:", result)
    return result
}
