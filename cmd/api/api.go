package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nakshatraraghav/notesgen/lib"
)

var log = lib.GetLogger()

type APIServer struct {
	addr string
}

func NewAPIServer() *APIServer {
	return &APIServer{
		addr: ":3000",
	}
}

func (api *APIServer) StartServer() error {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Info().Msg("server started on localhost:3000")

	return http.ListenAndServe(api.addr, router)
}
