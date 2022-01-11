package main

import (
	"github.com/mikejeuga/go-exercises/add"
	"os"
)

func main()  {
	inputTxt, err := os.Open("input.txt")
	if err != nil {
		return
	}
	add.Add(inputTxt, os.Stdout)
}





