package main

import (
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
	"github.com/mikejeuga/go-exercises/src/adapters/httpserver"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"log"
)

func main() {
	server := httpserver.NewServer(cli.AdderFunc(add.Add))
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}




