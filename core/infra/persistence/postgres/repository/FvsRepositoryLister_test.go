package pg_repository

import (
	"context"
	"fmt"
	"kaduhod/fin_v3/core/domain/investment"
	"kaduhod/fin_v3/core/infra/persistence/postgres/connection"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*func TestMain2(m *testing.M) {

    enviroment := os.Args[len(os.Args) - 1]
    envFile := ".env.development"
    if enviroment == "local" {
        envFile = ".env.local"
    }
    err := godotenv.Load("../../../../../" + envFile)
    if err != nil {
        log.Fatal("Erro ao carregar .env:", err)
    }

    os.Exit(m.Run())
}*/

func TestFvsResultRepositoryLister_Get(t *testing.T) {
    ctx := context.Background()
    conn := pg_connection.NewPgxConnection()
    defer conn.Conn.Close()

    repo := NewFvsPgRepositoryLister(conn)


    conn.Conn.Exec(ctx, `INSERT INTO fvs_results
(id, user_id, roi, roi_real, total_invested, initial_value, final_value, final_value_real, net_gain, net_gain_real, roi_porcentage, roi_porcentage_real, contribution, tax_real, tax, periods_json, periods_real_json, periods, tax_inflation, first_day, tags, created_at, deleted_at)
VALUES(19, 440, 250973.477194888, 136712.7889918184, 3243243.24, 3243243.24, 3494215.717194888, 3379956.0289918184, 3494215.717194888, 3379956.0289918184, 7.7383180545807, 4.21531099812971, 0, 0.0828625235404896, 0.15, '[{"date": "08/25", "period": 1, "accrued": "3.283.783,78", "interest": "40.540,54"}, {"date": "09/25", "period": 2, "accrued": "3.324.831,08", "interest": "41.047,30"}, {"date": "10/25", "period": 3, "accrued": "3.366.391,47", "interest": "41.560,39"}, {"date": "11/25", "period": 4, "accrued": "3.408.471,36", "interest": "42.079,89"}, {"date": "12/25", "period": 5, "accrued": "3.451.077,25", "interest": "42.605,89"}, {"date": "01/26", "period": 6, "accrued": "3.494.215,72", "interest": "43.138,47"}]'::jsonb, '[{"date": "08/25", "period": 1, "accrued": "3.265.638,52", "interest": "22.395,28"}, {"date": "09/25", "period": 2, "accrued": "3.288.188,44", "interest": "22.549,92"}, {"date": "10/25", "period": 3, "accrued": "3.310.894,07", "interest": "22.705,63"}, {"date": "11/25", "period": 4, "accrued": "3.333.756,49", "interest": "22.862,42"}, {"date": "12/25", "period": 5, "accrued": "3.356.776,78", "interest": "23.020,29"}, {"date": "01/26", "period": 6, "accrued": "3.379.956,03", "interest": "23.179,25"}]'::jsonb, 6, 0.062, true, NULL, '2025-07-08 00:04:27.365', NULL)`)
    conn.Conn.Exec(ctx, `INSERT INTO fvs_results
(id, user_id, roi, roi_real, total_invested, initial_value, final_value, final_value_real, net_gain, net_gain_real, roi_porcentage, roi_porcentage_real, contribution, tax_real, tax, periods_json, periods_real_json, periods, tax_inflation, first_day, tags, created_at, deleted_at)
VALUES(20, 440, 250974.477194888, 136712.7889918184, 3243243.24, 3243243.24, 3494215.717194888, 3379956.0289918184, 3494215.717194888, 3379956.0289918184, 7.7383180545807, 4.21531099812971, 0, 0.0828625235404896, 0.15, '[{"date": "08/25", "period": 1, "accrued": "3.283.783,78", "interest": "40.540,54"}, {"date": "09/25", "period": 2, "accrued": "3.324.831,08", "interest": "41.047,30"}, {"date": "10/25", "period": 3, "accrued": "3.366.391,47", "interest": "41.560,39"}, {"date": "11/25", "period": 4, "accrued": "3.408.471,36", "interest": "42.079,89"}, {"date": "12/25", "period": 5, "accrued": "3.451.077,25", "interest": "42.605,89"}, {"date": "01/26", "period": 6, "accrued": "3.494.215,72", "interest": "43.138,47"}]'::jsonb, '[{"date": "08/25", "period": 1, "accrued": "3.265.638,52", "interest": "22.395,28"}, {"date": "09/25", "period": 2, "accrued": "3.288.188,44", "interest": "22.549,92"}, {"date": "10/25", "period": 3, "accrued": "3.310.894,07", "interest": "22.705,63"}, {"date": "11/25", "period": 4, "accrued": "3.333.756,49", "interest": "22.862,42"}, {"date": "12/25", "period": 5, "accrued": "3.356.776,78", "interest": "23.020,29"}, {"date": "01/26", "period": 6, "accrued": "3.379.956,03", "interest": "23.179,25"}]'::jsonb, 6, 0.062, true, NULL, '2025-07-08 00:04:27.365', NULL)`)
    t.Run("Get investments by userId :: need success", func(t *testing.T) {
        results, err := repo.Get(investment.FutureValueOfASerieResult{Id: 0, UserId: 440})
        if err != nil {
            fmt.Println(err)
        }
        assert.NoError(t, err)
        assert.Greater(t, len(results), 1)
    })


}

