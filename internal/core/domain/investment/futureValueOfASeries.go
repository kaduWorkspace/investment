package investment

import valueobjects "kaduhod/fin_v3/internal/core/domain/valueObjects"
type FutureValueOfASeries interface {
    Calculate(contribuition, taxDecimal, periods valueobjects.Money, firstDay bool) float64
    monthlyTax(tax valueobjects.Money) valueobjects.Money
}

