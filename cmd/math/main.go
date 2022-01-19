package main

import (
	"github.com/mikejeuga/go-exercises/src/adapters/httpserver"
	"log"
)

func main() {
	server := httpserver.NewServer()
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}


}




