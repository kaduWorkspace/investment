package interface_chi

import (
	"kaduhod/fin_v3/core/domain/http"
	infra_investment "kaduhod/fin_v3/core/infra/investment/decimal"
	net_http "net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
type ServerChi struct {
    handler net_http.Handler
}
func NewServer() http.Server {
    return &ServerChi{}
}
func (s *ServerChi) Start(port string) {
    net_http.ListenAndServe(port, s.handler)
}
func (s *ServerChi) Setup() {
    compoundInterestServiceDecimal := infra_investment.CompoundInterestDecimal{}
    futureValueOfASeriesServiceDecimal := infra_investment.FutureValueOfASerieDecimal{}
    investmentHandler := InvestmentHandlerChi{
        CompoundInterestService: compoundInterestServiceDecimal,
        FutureValueOfASeriesService: futureValueOfASeriesServiceDecimal,
    }
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Get("/health-check", func(w net_http.ResponseWriter, r *net_http.Request) {
        w.Write([]byte("✅"))
    })
    r.Route("/api/investments", func(r chi.Router) {
        r.Post("/compound-interest", investmentHandler.CompoundInterestApi)
        r.Post("/future-value-of-a-series", investmentHandler.FutureValueOfASeriesApi)
    })
    r.Route("/web/investments", func(r chi.Router) {
        r.Post("/compound-interest", investmentHandler.CompoundInterest)
        r.Post("/future-value-of-a-series", investmentHandler.FutureValueOfASeries)
    })
    s.handler = r
}
