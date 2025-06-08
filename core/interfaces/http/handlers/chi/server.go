package interface_chi

import (
	domain_http "kaduhod/fin_v3/core/domain/http"
	auth_std "kaduhod/fin_v3/core/infra/auth/std"
	infra_investment "kaduhod/fin_v3/core/infra/investment/decimal"
	"kaduhod/fin_v3/core/infra/session/memory"
	http_middleware "kaduhod/fin_v3/core/interfaces/http/middlewares/http"
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
    err := http.ListenAndServe(port, s.handler)
    if err != nil {
        panic(err)
    }
}
func (s *ServerChi) Setup() {
    compoundInterestServiceDecimal := infra_investment.CompoundInterestDecimal{}
    futureValueOfASeriesServiceDecimal := infra_investment.FutureValueOfASerieDecimal{}
    rootDir, _ := os.Getwd()
    rndr, err := renderer.NewRenderer(rootDir+"/core/interfaces/web/components", rootDir+"/core/interfaces/web/pages")
    if err != nil {
        panic(err)
    }
    investmentHandler := NewInvestmentHandler(compoundInterestServiceDecimal, futureValueOfASeriesServiceDecimal)
    inMemorySessionService := memory.NewInMemorySession()
    investmentHandlerWeb := NewInvestmentHandlerChiWeb(inMemorySessionService ,compoundInterestServiceDecimal, futureValueOfASeriesServiceDecimal, rndr)
    sessionMidlewareHandler := http_middleware.NewSessionHandlerMiddleware(inMemorySessionService)
    csrfMiddlewareHandler := http_middleware.NewCsrfHandlerMiddleware(inMemorySessionService)
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Handle("/public/*", http.StripPrefix("/public/" ,http.FileServer(http.Dir(rootDir + "/public"))))
    r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("ok"))
    })
    r.Route("/api", func(r chi.Router) {
        r.Use(auth_std.AuthTokenMiddleware)
        r.Route("/investments", func(r chi.Router) {
            r.Post("/compound-interest", investmentHandler.CompoundInterestApi)
            r.Post("/future-value-of-a-series", investmentHandler.FutureValueOfASeriesWithTrackingApi)
            r.Post("/future-value-of-a-series-simple", investmentHandler.FutureValueOfASeriesApi)
            r.Post("/future-value-of-a-series/predict-contribution", investmentHandler.PredictFV)
        })
    })
    r.Route("/web", func(r chi.Router) {
        r.Use(sessionMidlewareHandler.CheckSessionMiddleware)
        r.Route("/investments", func(r chi.Router) {
            r.Get("/fv", investmentHandlerWeb.FutureValueOfASeriesFormPage)
            r.With(csrfMiddlewareHandler.ValidateCsrfMiddleware).Post("/fv", investmentHandlerWeb.FutureValueOfASeriesResultPage)
            r.Get("/fv/predict", investmentHandlerWeb.FutureValueOfASeriesPredictFormPage)
            r.With(csrfMiddlewareHandler.ValidateCsrfMiddleware).Post("/fv/predict", investmentHandlerWeb.FutureValueOfASeriesPredictResultPage)
        })
    })
    r.With(sessionMidlewareHandler.CreateSessionMiddleware).Get("/", investmentHandlerWeb.Index)
    s.handler = r
}
