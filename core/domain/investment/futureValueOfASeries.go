package investment

import (
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	"time"
)
type PeriodTracker struct {
    Accrued valueobjects.Money `json:"accrued"`
    Period  int `json:"period"`
    Interest valueobjects.Money `json:"interest"`
    Date time.Time `json:"date"`
    InterestFormated string `json:"interest_formated"`
    AccruedFormated string `json:"accrued_formated"`
    DateFormated string `json:"date_formated"`
    Acumulated string `json:"acumulated"`
}
func NewPeriodTracker(accrued valueobjects.Money, period int, interest valueobjects.Money, initialDate time.Time) PeriodTracker {
    return PeriodTracker{
        Accrued: accrued,
        Period: period,
        Interest: interest,
        Date: initialDate,
        InterestFormated: interest.Formatted(),
        AccruedFormated: accrued.Formatted(),
        Acumulated: accrued.Formatted(),
        DateFormated: GetDateFormated(initialDate.AddDate(0, period, 0)),
    }
}
func GetDateFormated(date time.Time) string {
    return date.Format("02/01/2006")
}
type FutureValueOfASeries interface {
    Calculate(contribuition, taxDecimal, periods valueobjects.Money, firstDay bool) float64
    monthlyTax(tax valueobjects.Money) valueobjects.Money
    CalculateTrackingPeriods(initialValue, contribuition, taxDecimal, periods valueobjects.Money, firstDay bool, initialDate time.Time) (valueobjects.Money, []PeriodTracker)
}

