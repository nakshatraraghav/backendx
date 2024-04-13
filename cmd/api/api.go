package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nakshatraraghav/notesgen/lib"
	"github.com/nakshatraraghav/notesgen/middlewares"
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

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Info().Msg("server started on localhost" + lib.ENV.Addr)

	return http.ListenAndServe(api.addr, router)
}

func (api *APIServer) atttachMiddlewares(router *chi.Mux) {
	router.Use(middlewares.LogRequests)
}
