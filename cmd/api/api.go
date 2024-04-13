package api

import (
	"fmt"
	"net/http"
)

type APIServer struct {
	addr string
}

func NewAPIServer() *APIServer {
	return &APIServer{
		addr: ":3000",
	}
}

func (api *APIServer) StartServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	fmt.Println("server started on localhost:3000")

	return http.ListenAndServe(api.addr, nil)
}
