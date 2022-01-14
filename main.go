package main

import (
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/mikejeuga/go-exercises/src/adapters/cli"
	"github.com/spf13/pflag"
	"os"
	"strconv"
)

func main()  {
	cmd := cli.NewCommand()

	var intCheck []int
	for _, val := range os.Args[1:] {
	num, _ := strconv.Atoi(val)
	intCheck = append(intCheck, num)
}
	if len(os.Args[1:]) > 0 && cmd.Service.Add(intCheck...) > 0 {
		fmt.Fprintf(os.Stdout, "%s\n", humanize.Commaf(float64(cmd.Service.Add(intCheck...))))
		return
	}



values:= []string{"input.csv"}
stringP := pflag.StringSliceP("input-file", "f", values,"Enter the name of the files to be processed")
pflag.Parse()


for _, file := range *stringP {
	readFile, _ := os.ReadFile("data/"+file)
	defaultFile := bytes.NewReader(readFile)
	cmd.Add(defaultFile, os.Stdout)
	}
}
