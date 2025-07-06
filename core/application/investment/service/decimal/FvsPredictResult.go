package app_investment_decimal

import (
	"context"
	"fmt"
	"kaduhod/fin_v3/core/domain/investment"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
)

type FVSPredictResultService struct {
    connection *pg_connection.PgxConnextion
}
func NewFVSPredictResult(conn *pg_connection.PgxConnextion) investment.InvestmentResultService[investment.FutureValueOfASeriePredictResult] {
    return &FVSPredictResultService{connection: conn}
}

func (s *FVSPredictResultService) Save(result *investment.FutureValueOfASeriePredictResult) error {
    exists, err := s.CheckIfAlreadyExists(result)
    if err != nil {
        return err
    }
    if exists {
        return nil
    }
	query := `
		INSERT INTO fvs_predict_results (
			user_id,
			initial_value,
			final_value,
			contribution,
			tax_real,
			tax,
			periods,
			tax_inflation,
			first_day
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
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
		result.InitialValue,
		result.FinalValue,
		result.Contribution,
		result.TaxReal,
		result.Tax,
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
func (s *FVSPredictResultService) CheckIfAlreadyExists(result *investment.FutureValueOfASeriePredictResult) (bool, error) {
    query := `
        SELECT EXISTS (
            SELECT 1 FROM fvs_predict_results
            WHERE
                initial_value = $1 AND
                user_id = $2 AND
                contribution = $3 AND
                tax = $4 AND
                tax_inflation = $5 AND
                periods = $6 AND
                first_day = $7
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
        result.FirstDay,
    ).Scan(&exists)
    return exists, err
}
