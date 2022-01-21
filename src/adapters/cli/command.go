package cli

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/dustin/go-humanize"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"github.com/spf13/pflag"
	"io"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Calc add.Adder
}

func NewCommand(calc add.Adder) *Command {
	return &Command{Calc: calc}
}


func (c Command) Add(ctx context.Context, w io.Writer, r io.Reader) {
	readAll, err := io.ReadAll(r)
	if err != nil {
		panic("it doesn't read io.ReadAll")
	}

	values := getData(readAll)
	toInts := stringsToInts(values)
	for _, num := range toInts {
		fmt.Fprintf(w, "%v\n", num)
	}
	sum := c.Calc.Add(toInts...)
	fmt.Fprintf(w, "Total: %s\n", humanize.Commaf(float64(sum)))
}

func (c *Command)AddWithFileInputs(ctx context.Context) {
	values := []string{"data/input.csv"}
	stringP := pflag.StringSliceP("input-file", "f", values, "Enter the name of the files to be processed")
	pflag.Parse()

	for _, file := range *stringP {
		readFile, _ := os.ReadFile(file)
		defaultFile := bytes.NewReader(readFile)
		c.Add(ctx, os.Stdout, defaultFile)
	}
}

func (c *Command) AddManualEntries(ctx context.Context) bool {
	var intCheck []int
	for _, val := range os.Args[1:] {
		num, _ := strconv.Atoi(val)
		intCheck = append(intCheck, num)
	}

	if len(os.Args[1:]) > 0 && c.Calc.Add(intCheck...) > 0 {
		for _, val := range intCheck {
			fmt.Println(val)
		}
		fmt.Fprintf(os.Stdout, "Total: %s\n", humanize.Commaf(float64(c.Calc.Add(intCheck...))))
		return true
	}
	return false
}


func stringsToInts(numbers []string) []int {
	var values []int
	for _, val := range numbers {
		num, err := strconv.Atoi(val)
		if err != nil {
			continue
		}
		values = append(values, num)
	}
	return values
}



func getData(in []byte) []string{
	if len(dataSpaceDelimited(in)) > 2 {
		return dataSpaceDelimited(in)
	}

	if len(dataNewLine(in)) > len(dataCSV(in)) {
		return dataNewLine(in)
	}
	return dataCSV(in)
}


func dataNewLine(in []byte) []string {
	reader := bytes.NewReader(in)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	var numbers []string
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	return numbers
}

func dataSpaceDelimited(in []byte) []string{
	reader := bytes.NewReader(in)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := scanner.Text()

	split := strings.Split(input, " ")
	return split
}


func dataCSV(in []byte) []string{
	reader := bytes.NewReader(in)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	input := scanner.Text()

	split := strings.Split(input, ",")
	return split
}
