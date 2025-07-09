package pg_repository

import (
	"context"
	"errors"
	"fmt"
	"kaduhod/fin_v3/core/domain/investment"
	"kaduhod/fin_v3/core/domain/repository"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	"strings"
)

type FvsRepositoryLister struct {
    conn *pg_connection.PgxConnextion
}

func NewFvsPgRepositoryLister(conn *pg_connection.PgxConnextion) repository.RepositoryList[investment.FutureValueOfASerieResult] {
    return &FvsRepositoryLister{conn: conn}
}
func (r *FvsRepositoryLister) Get(filters investment.FutureValueOfASerieResult) ([]investment.FutureValueOfASerieResult, error) {
    ctx := context.Background()
    var investments []investment.FutureValueOfASerieResult
    // Build WHERE clause and arguments
    var whereClause strings.Builder
    var args []interface{}
    var conditions []string

    if filters.Id != 0 {
        args = append(args, filters.Id)
        conditions = append(conditions, "id = $1")
    } else {
        cont := 1
        if filters.UserId != 0 {
            args = append(args, filters.UserId)
            conditions = append(conditions, fmt.Sprintf("user_id = $%d", cont))
            cont++
        }
        /*if filters.Email != "" {
            args = append(args, filters.Email)
            conditions = append(conditions, fmt.Sprintf("email = $%d", cont))
            cont++
        }*/
    }
    if len(args) == 0 {
        return investments, errors.New("no filter criteria provided")
    }

    whereClause.WriteString("WHERE ")
    whereClause.WriteString(strings.Join(conditions, " AND "))

    query := `
    SELECT
        id,
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
        first_day,
        tags,
        created_at,
        deleted_at
    FROM fvs_results
    ` + whereClause.String()
    rows, err := r.conn.Conn.Query(ctx, query, args...)
    defer rows.Close()
    if err != nil {
        return investments, fmt.Errorf("failed to get investments: %w", err)
    }
    for rows.Next() {
        var row investment.FutureValueOfASerieResult

        err := rows.Scan(
            &row.Id,
            &row.UserId,
            &row.ROI,
            &row.ROIReal,
            &row.TotalInvested,
            &row.InitialValue,
            &row.FinalValue,
            &row.FinalValueReal,
            &row.NetGain,
            &row.NetGainReal,
            &row.ROIPorcentage,
            &row.ROIPorcentageReal,
            &row.Contribution,
            &row.TaxReal,
            &row.Tax,
            &row.PeriodsJSON,
            &row.PeriodsRealJSON,
            &row.Periods,
            &row.TaxInflation,
            &row.FirstDay,
            &row.Tags,
            &row.CreatedAt,
            &row.DeletedAt,
        )
        if err != nil {
            fmt.Println("Error scanning row:", err)
            return nil, err
        }

        investments = append(investments, row)
    }

    return investments, nil
}
