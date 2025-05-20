package infra_investment

import (
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
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
