package app_investment_decimal

import (
	"context"
	"fmt"
	"kaduhod/fin_v3/core/domain/investment"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
)

type InvestmentResultServicePg struct {
    connection *pg_connection.PgxConnextion
}
func NewFVSResult(conn *pg_connection.PgxConnextion) investment.InvestmentResultService[investment.FutureValueOfASerieResult] {
    return &InvestmentResultServicePg{connection: conn}
}

func (s *InvestmentResultServicePg) Save(result *investment.FutureValueOfASerieResult) error {
    exists, err := s.CheckIfAlreadyExists(result)
    if err != nil {
        return err
    }
    if exists {
        return nil
    }
	query := `
		INSERT INTO investment_results (
			user_id,
			roi,
			roi_real,
			total_invested,
			initial_value,
			final_value,
			final_value_real,
			net_gain,
			net_gain_real,
			roi_porcentage,
			roi_porcentage_real,
			contribution,
			tax_real,
			tax,
			periods_json,
			periods_real_json,
			periods,
			tax_inflation,
			first_day
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15, $16, $17, $18,
			$19
		)
	`
    ctx := context.Background()
    tx, err := s.connection.Conn.Begin(ctx)
    if err != nil {
        fmt.Println(err)
        return err
    }
    _, err = tx.Exec(
        ctx,
		query,
		result.UserID,
		result.ROI,
		result.ROIReal,
		result.TotalInvested,
		result.InitialValue,
		result.FinalValue,
		result.FinalValueReal,
		result.NetGain,
		result.NetGainReal,
		result.ROIPorcentage,
		result.ROIPorcentageReal,
		result.Contribution,
		result.TaxReal,
		result.Tax,
		result.PeriodsJSON,
		result.PeriodsRealJSON,
        result.Periods,
		result.TaxInflation,
        result.FirstDay,
	)
    if err != nil {
        fmt.Println(err)
        tx.Rollback(ctx)
        return err
    }
    return tx.Commit(ctx)
}
func (s *InvestmentResultServicePg) CheckIfAlreadyExists(result *investment.FutureValueOfASerieResult) (bool, error) {
    query := `
        SELECT EXISTS (
            SELECT 1 FROM investment_results
            WHERE
                initial_value = $1 AND
                user_id = $2 AND
                contribution = $3 AND
                tax = $4 AND
                tax_inflation = $5 AND
                periods = $6
        )`
    ctx := context.Background()
    var exists bool
    err := s.connection.Conn.QueryRow(ctx, query,
        result.InitialValue,
        result.UserID,
        result.Contribution,
        result.Tax,
        result.TaxInflation,
        result.Periods,
    ).Scan(&exists)
    return exists, err
}
