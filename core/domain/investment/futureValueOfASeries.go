package investment

import (
	valueobjects "kaduhod/fin_v3/core/domain/valueObjects"
	"time"
)
type PeriodTracker struct {
    Accrued valueobjects.Money `json:"-"`
    Period  int `json:"period"`
    Interest valueobjects.Money `json:"-"`
    Date time.Time `json:"-"`
    InterestFormated string `json:"interest"`
    AccruedFormated string `json:"accrued"`
    DateFormated string `json:"date"`
}
func NewPeriodTracker(accrued valueobjects.Money, period int, interest valueobjects.Money, initialDate time.Time) PeriodTracker {
    return PeriodTracker{
        Accrued: accrued,
        Period: period,
        Interest: interest,
        Date: initialDate,
        InterestFormated: interest.Formatted(),
        AccruedFormated: accrued.Formatted(),
        DateFormated: GetDateFormated(initialDate.AddDate(0, period, 0)),
    }
}
func GetDateFormated(date time.Time) string {
    return date.Format("02/01/2006")
}
type FutureValueOfASeries interface {
    Calculate(contribution, taxDecimal valueobjects.Money, firstDay bool, periods int) valueobjects.Money
    CalculateTrackingPeriods(initialValue, contribution, taxDecimal valueobjects.Money, firstDay bool, initialDate time.Time, periods int) (valueobjects.Money, []PeriodTracker)
    PredictContribution(finalValue, taxDecimal, initialValue valueobjects.Money, contributionOnFirstDay bool, periods int) (valueobjects.Money)
}

