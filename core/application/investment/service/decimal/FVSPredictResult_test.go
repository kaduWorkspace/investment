
package app_investment_decimal

import (
	"context"
	"log"
	"os"
	"testing"

	"kaduhod/fin_v3/core/domain/investment"
	pg_connection "kaduhod/fin_v3/core/infra/persistence/postgres/connection"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)
func TestInvestmentResultPredictServicePg_Save(t *testing.T) {
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

	service := NewFVSPredictResult(conn)

	// Garante limpeza dos dados com o mesmo user_id para não afetar outros testes
	userID := 1218
	cleanup := func() {
		conn.Conn.Exec(ctx, "DELETE FROM fvs_predict_results WHERE user_id = $1", userID)
	}
	cleanup() // Limpa antes
	defer cleanup()

	result := &investment.FutureValueOfASeriePredictResult{
		UserID:             userID,
		InitialValue:       toSqlNullFloat(8000.0),
		FinalValue:         toSqlNullFloat(11500.0),
		Contribution:       toSqlNullFloat(2000.0),
		TaxReal:            toSqlNullFloat(0.05),
		Tax:                toSqlNullFloat(0.1),
		Periods:            12,
		TaxInflation:       toSqlNullFloat(0.03),
		FirstDay:           true,
	}

    err = service.Save(result)
	assert.NoError(t, err)

	// Verifica se foi salvo
	var count int
	err = conn.Conn.QueryRow(ctx, `
		SELECT COUNT(*) FROM fvs_predict_results
		WHERE user_id = $1 AND initial_value = $2 AND contribution = $3
	`, result.UserID, result.InitialValue, result.Contribution).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)


    exists, err := service.CheckIfAlreadyExists(result)
    assert.NoError(t, err)
    assert.True(t, exists)
}
