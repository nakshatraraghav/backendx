package main

import (
	"log"

	"github.com/nakshatraraghav/notesgen/cmd/api"
)

func main() {
	server := api.NewAPIServer()

	err := server.StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
