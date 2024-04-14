package main

import (
	"github.com/nakshatraraghav/backendx/cmd/api"
	"github.com/nakshatraraghav/backendx/lib"
)

var log = lib.GetLogger()

func main() {
	server := api.NewAPIServer()

	err := server.StartServer()
	if err != nil {
		log.Fatal().Err(err)
	}
}
