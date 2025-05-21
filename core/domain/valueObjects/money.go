package valueobjects

type Money interface {
    Add(Money) Money
    Subtract(Money) Money
    Multiply(Money) Money
    Divide(Money) Money
    Equals(Money) bool
    Pow(Money) Money
    GetAmount() float64
    Formatted() string
}
