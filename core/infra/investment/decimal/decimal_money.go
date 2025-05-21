package infra_investment

import (
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

type DecimalMoney struct {
    amount float64
}
func NewDecimalMoney(amount float64) valueobjects.Money {
    return DecimalMoney{amount}
}
func (m DecimalMoney) GetAmount() float64 {
    return m.amount
}
func (m DecimalMoney) Formatted() string {
    str := strconv.FormatFloat(m.GetAmount(), 'f', 2, 64)
    str = strings.Replace(str, ".", ",", 1)
    parts := strings.Split(str, ",")
    intPart := parts[0]
    decimalPart := parts[1]
    var result string
    count := 0
    for i := len(intPart) - 1; i >= 0; i-- {
        result = string(intPart[i]) + result
        count++
        if count%3 == 0 && i > 0 {
            result = "." + result
        }
    }
    return result + "," + decimalPart
}
func (m DecimalMoney) Pow(power valueobjects.Money) valueobjects.Money {
    decimal_a := decimal.NewFromFloat(m.GetAmount())
    decimal_b := decimal.NewFromFloat(power.GetAmount())
    result, _ := decimal_a.Pow(decimal_b).Round(16).Float64()
    return NewDecimalMoney(result)
}
func (m DecimalMoney) Add(add valueobjects.Money) valueobjects.Money {
    decimal_b := decimal.NewFromFloat(add.GetAmount())
    decimal_a := decimal.NewFromFloat(m.GetAmount())
    result := decimal_a.Add(decimal_b)
    result_rounded, _ := result.Round(16).Float64()
    return NewDecimalMoney(result_rounded)
}
func (m DecimalMoney) Subtract(sub valueobjects.Money) valueobjects.Money {
    decimal_b := decimal.NewFromFloat(sub.GetAmount())
    decimal_a := decimal.NewFromFloat(m.GetAmount())
    result := decimal_a.Sub(decimal_b)
    result_rounded, _ := result.Round(16).Float64()
    return NewDecimalMoney(result_rounded)
}
func (m DecimalMoney) Multiply(multiplier valueobjects.Money) valueobjects.Money {
    decimal_b := decimal.NewFromFloat(multiplier.GetAmount())
    decimal_a := decimal.NewFromFloat(m.GetAmount())
    result := decimal_a.Mul(decimal_b)
    result_rounded, _ := result.Round(16).Float64()
    return NewDecimalMoney(result_rounded)
}
func (m DecimalMoney) Divide(divisor valueobjects.Money) valueobjects.Money {
    decimal_b := decimal.NewFromFloat(divisor.GetAmount())
    decimal_a := decimal.NewFromFloat(m.GetAmount())
    result := decimal_a.Div(decimal_b)
    result_rounded, _ := result.Round(16).Float64()
    return NewDecimalMoney(result_rounded)
}
func (m DecimalMoney) Equals(other valueobjects.Money) bool {
    return m.GetAmount() == other.GetAmount()
}
