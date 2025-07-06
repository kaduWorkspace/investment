package app_investment_decimal

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"kaduhod/fin_v3/core/domain/investment"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)
func toSqlNullFloat(n float64) sql.NullFloat64 {
    return sql.NullFloat64{Float64: n, Valid: true}
}
func TestInvestmentResultServicePg_Save(t *testing.T) {
    enviroment := os.Args[len(os.Args) - 1]
    envFile := ".env.development"
    if enviroment == "local" {
        envFile = "/home/carlos/projetos/meu-app/chi_version/.env.local"
    } else {
       envFile = "../../../../../" + envFile
    }
    err := godotenv.Load(envFile)
    if err != nil {
        log.Fatal("Erro ao carregar .env:", err)
    }


	ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
	defer conn.Conn.Close()

	service := NewInvestmentResultServicePg(conn)

	// Garante limpeza dos dados com o mesmo user_id para não afetar outros testes
	userID := 1218
	cleanup := func() {
		conn.Conn.Exec(ctx, "DELETE FROM investment_results WHERE user_id = $1", userID)
	}
	cleanup() // Limpa antes
	defer cleanup()

	result := &investment.InvestmentResult{
		UserID:             userID,
		ROI:                toSqlNullFloat(0.15),
		ROIReal:            toSqlNullFloat(0.12),
		TotalInvested:      toSqlNullFloat(10000.0),
		InitialValue:       toSqlNullFloat(8000.0),
		FinalValue:         toSqlNullFloat(11500.0),
		FinalValueReal:     toSqlNullFloat(11200.0),
		NetGain:            toSqlNullFloat(3500.0),
		NetGainReal:        toSqlNullFloat(3200.0),
		ROIPorcentage:      toSqlNullFloat(0.4375),
		ROIPorcentageReal:  toSqlNullFloat(0.4),
		Contribution:       toSqlNullFloat(2000.0),
		TaxReal:            toSqlNullFloat(0.05),
		Tax:                toSqlNullFloat(0.1),
		PeriodsJSON:        []byte(`{"example": "test"}`),
		PeriodsRealJSON:    []byte(`{"real": "json"}`),
		Periods:            12,
		TaxInflation:       toSqlNullFloat(0.03),
		FirstDay:           true,
	}

    err = service.Save(result)
	assert.NoError(t, err)

	// Verifica se foi salvo
	var count int
	err = conn.Conn.QueryRow(ctx, `
		SELECT COUNT(*) FROM investment_results
		WHERE user_id = $1 AND initial_value = $2 AND contribution = $3
	`, result.UserID, result.InitialValue, result.Contribution).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)


    exists, err := service.CheckIfAlreadyExists(result)
    assert.NoError(t, err)
    assert.True(t, exists)
}

