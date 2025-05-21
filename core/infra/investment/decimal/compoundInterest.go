package infra_investment

import (
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
)

type CompoundInterestDecimal struct {}
func (self CompoundInterestDecimal) Calculate(initialValue, taxDecimal valueobjects.Money, periods int) float64 {
    one := NewDecimalMoney(1.0)
    twelve := NewDecimalMoney(12.0)
    taxOverTwelve := taxDecimal.Divide(twelve)
    onePlusTax := one.Add(taxOverTwelve)
    powOnePlusTax := onePlusTax.Pow(NewDecimalMoney(float64(periods)))
    multipliedValue := initialValue.Multiply(powOnePlusTax)
    return multipliedValue.GetAmount()
}

