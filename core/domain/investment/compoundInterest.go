package investment

import valueobjects "kaduhod/fin_v3/core/domain/valueObjects"


type CompoundInterest interface {
    Calculate(initialValue, taxDecimal valueobjects.Money, periods int) valueobjects.Money
    CalculateRealValue(initialValue, taxDecimal, inflationTax valueobjects.Money, periods int) valueobjects.Money
}


