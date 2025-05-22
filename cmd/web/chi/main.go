package chi_web

import (
	infra_investment "kaduhod/fin_v3/core/infra/investment/decimal"
	interface_chi "kaduhod/fin_v3/core/interfaces/http/handlers/chi"

	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    compoundInterestServiceDecimal := infra_investment.CompoundInterestDecimal{}
    futureValueOfASeriesServiceDecimal := infra_investment.FutureValueOfASerieDecimal{}

    investmentHandler := interface_chi.InvestmentHandlerChi{
        CompoundInterestService: compoundInterestServiceDecimal,
        FutureValueOfASeriesService: futureValueOfASeriesServiceDecimal,
    }

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Route("/api/investments", func(r chi.Router) {
        r.Post("/compound-interest", investmentHandler.CompoundInterestApi)
        r.Post("/future-value-of-a-series", investmentHandler.FutureValueOfASeriesApi)
    })
}
