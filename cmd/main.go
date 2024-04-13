package main

import (
	"github.com/nakshatraraghav/notesgen/cmd/api"
	"github.com/nakshatraraghav/notesgen/lib"
)

var log = lib.GetLogger()

func main() {
	server := api.NewAPIServer()

	err := server.StartServer()
	if err != nil {
		log.Fatal().Err(err)
	}
}
