package investment

import valueobjects "kaduhod/fin_v3/internal/core/domain/valueObjects"


type CompoundInterest interface {
    Calculate(initialValue, taxDecimal, periods valueobjects.Money) float64
}


