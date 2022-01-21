package main

import (
	"github.com/mikejeuga/go-exercises/src/adapters/httpserver"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"log"
)

func main() {
	server := httpserver.NewServer(add.Default)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}




