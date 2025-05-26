package interface_chi

import (
	domain_http "kaduhod/fin_v3/core/domain/http"
	auth_std "kaduhod/fin_v3/core/infra/auth/std"
	infra_investment "kaduhod/fin_v3/core/infra/investment/decimal"
	"kaduhod/fin_v3/core/interfaces/web/renderer"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
type ServerChi struct {
    handler http.Handler
}
func NewServer() domain_http.Server {
    return &ServerChi{}
}

func (s *ServerChi) Start(port string) {
    http.ListenAndServe(port, s.handler)
}
func (s *ServerChi) Setup() {
    compoundInterestServiceDecimal := infra_investment.CompoundInterestDecimal{}
    futureValueOfASeriesServiceDecimal := infra_investment.FutureValueOfASerieDecimal{}
    rootDir, _ := os.Getwd()
    rndr, err := renderer.NewRenderer(rootDir+"/core/interfaces/web/components", rootDir+"/core/interfaces/web/pages")
    if err != nil {
        panic(err)
    }
    investmentHandler := InvestmentHandlerChi{
        CompoundInterestService: compoundInterestServiceDecimal,
        FutureValueOfASeriesService: futureValueOfASeriesServiceDecimal,
        Renderer: rndr,
    }
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Handle("/public/*", http.StripPrefix("/public/" ,http.FileServer(http.Dir(rootDir + "/public"))))
    r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })
    r.Route("/api/investments", func(r chi.Router) {
        r.Use(auth_std.AuthTokenMiddleware)
        r.Post("/compound-interest", investmentHandler.CompoundInterestApi)
        r.Post("/future-value-of-a-series", investmentHandler.FutureValueOfASeriesWithTrackingApi)
        r.Post("/future-value-of-a-series-simple", investmentHandler.FutureValueOfASeriesApi)
        r.Post("/future-value-of-a-series/predict-contribution", investmentHandler.PredictFV)
    })
    r.Route("/web/investments", func(r chi.Router) {
        r.Get("/fv", investmentHandler.FutureValueOfASeriesForm)
        r.Get("/fv/predict", investmentHandler.FutureValueOfASeriesPredictForm)
        r.Post("/compound-interest", investmentHandler.CompoundInterest)
        r.Post("/future-value-of-a-series", investmentHandler.FutureValueOfASeries)
    })
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        if err := rndr.Render(w, "base", nil); err != nil {
            w.WriteHeader(500)
        }
    })
    s.handler = r
}
