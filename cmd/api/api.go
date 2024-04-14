package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nakshatraraghav/backendx/lib"
	"github.com/nakshatraraghav/backendx/middlewares"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var log = lib.GetLogger()

type APIServer struct {
	addr string
}

func initialize() {
	lib.LoadEnv()
}

func NewAPIServer() *APIServer {
	initialize()

	return &APIServer{
		addr: lib.ENV.Addr,
	}
}

func (api *APIServer) StartServer() error {
	router := chi.NewRouter()

	api.atttachMiddlewares(router)
	api.serverPrometheusMetrics(router)

	log.Info().Msg("server started on localhost" + lib.ENV.Addr)

	return http.ListenAndServe(api.addr, router)
}

func (api *APIServer) atttachMiddlewares(router *chi.Mux) {
	router.Use(middlewares.LogRequests)
}

func (api *APIServer) serverPrometheusMetrics(router *chi.Mux) {
	router.Get("/metrics", promhttp.Handler().ServeHTTP)
}
