package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nakshatraraghav/notesgen/lib"
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

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println(lib.ENV.Addr)

	log.Info().Msg("server started on localhost" + lib.ENV.Addr)

	return http.ListenAndServe(api.addr, router)
}
