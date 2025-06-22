package app_investment_decimal

import (
	"kaduhod/fin_v3/core/domain/investment"
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
)

type CompoundInterestDecimal struct {}
func (self CompoundInterestDecimal) Calculate(initialValue, taxDecimal valueobjects.Money, periods int) valueobjects.Money {
    one := NewDecimalMoney(1.0)
    twelve := NewDecimalMoney(12.0)
    taxOverTwelve := taxDecimal.Divide(twelve)
    onePlusTax := one.Add(taxOverTwelve)
    powOnePlusTax := onePlusTax.Pow(NewDecimalMoney(float64(periods)))
    multipliedValue := initialValue.Multiply(powOnePlusTax)
    return multipliedValue
}
func (self CompoundInterestDecimal) CalculateRealValue(initialValue, taxDecimal, inflationTax valueobjects.Money, periods int) valueobjects.Money {
    one := NewDecimalMoney(1.0)
    realTax := one.Add(taxDecimal).Divide(one.Add(inflationTax)).Subtract(one)
    return self.Calculate(initialValue, realTax, periods)
}
func NewCompoundInterest() investment.CompoundInterest {
    return &CompoundInterestDecimal{}
}

