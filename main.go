package main

import (
	"github.com/mikejeuga/go-exercises/add"
	"os"
)

func main()  {
	inputTxt, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	add.Add(inputTxt, os.Stdout)
}





