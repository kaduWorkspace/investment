package investment

import (
	"database/sql"
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
    return date.Format("01/06")
}
type FutureValueOfASeries interface {
    Calculate(contribution, taxDecimal valueobjects.Money, firstDay bool, periods int) valueobjects.Money
    CalculateRealValue(contribution, taxDecimal, taxInflation valueobjects.Money, firstDay bool, periods int) valueobjects.Money
    CalculateTrackingPeriods(initialValue, contribution, taxDecimal valueobjects.Money, firstDay bool, initialDate time.Time, periods int) (valueobjects.Money, []PeriodTracker)
    CalculateTrackingPeriodsRealValue(initialValue, contribution, taxDecimal, taxInflation valueobjects.Money, firstDay bool, initialDate time.Time, periods int) (valueobjects.Money, []PeriodTracker)
    PredictContribution(finalValue, taxDecimal, initialValue valueobjects.Money, contributionOnFirstDay bool, periods int) (valueobjects.Money)
    PredictContributionRealValue(finalValue, taxDecimal, initialValue, taxInflation valueobjects.Money, contributionOnFirstDay bool, periods int) (valueobjects.Money)
}
type FutureValueOfASerieResult struct {
	Id                   int             `json:"id"`
	UserId               int             `json:"user_id"`
	ROI                  sql.NullFloat64 `json:"roi"`
	ROIReal              sql.NullFloat64 `json:"roi_real"`
	TotalInvested        sql.NullFloat64 `json:"total_invested"`
	InitialValue         sql.NullFloat64 `json:"initial_value"`
	FinalValue           sql.NullFloat64 `json:"final_value"`
	FinalValueReal       sql.NullFloat64 `json:"final_value_real"`
	NetGain              sql.NullFloat64 `json:"net_gain"`
	NetGainReal          sql.NullFloat64 `json:"net_gain_real"`
	ROIPorcentage        sql.NullFloat64 `json:"roi_porcentage"`
	ROIPorcentageReal    sql.NullFloat64 `json:"roi_porcentage_real"`
	Contribution         sql.NullFloat64 `json:"contribution"`
	TaxReal              sql.NullFloat64 `json:"tax_real"`
	Tax                  sql.NullFloat64 `json:"tax"`
	PeriodsJSON          []byte          `json:"periods_json"`
	PeriodsRealJSON      []byte          `json:"periods_real_json"`
	Periods             int             `json:"periods"`
	TaxInflation        sql.NullFloat64 `json:"tax_inflation"`
	FirstDay            bool            `json:"first_day"`
	Tags []string `json:"tags"`
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}
type FutureValueOfASeriePredictResult struct {
	ID                   int             `json:"id"`
	UserID               int             `json:"user_id"`
	InitialValue         sql.NullFloat64 `json:"initial_value"`
	FinalValue           sql.NullFloat64 `json:"final_value"`
	Contribution         sql.NullFloat64 `json:"contribution"`
	TaxReal              sql.NullFloat64 `json:"tax_real"`
	Tax                  sql.NullFloat64 `json:"tax"`
	Periods              int             `json:"periods"`
	TaxInflation         sql.NullFloat64 `json:"tax_inflation"`
	FirstDay             bool            `json:"first_day"`
	Tags                 []string `json:"tags"`
	CreatedAt            time.Time    `json:"created_at"`
	DeletedAt            sql.NullTime `json:"deleted_at"`
}
type InvestmentResultService[T any] interface {
    Save(result *T) error
    CheckIfAlreadyExists(result *T) (bool, error)
}

