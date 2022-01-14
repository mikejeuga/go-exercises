package cli

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dustin/go-humanize"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"io"
	"strconv"
	"strings"
)

type Command struct {
	Calc add.Adder
}


func NewCommand() *Command {
	return &Command{add.NewAddProvider()}
}

func (c Command) Add(r io.Reader, w io.Writer) {
	readAll, err := io.ReadAll(r)
	if err != nil {
		panic("it doesn't read io.ReadAll")
	}

	values := getData(readAll)
	toInts := stringsToInts(values)
	sum := c.Calc.Add(toInts...)
	fmt.Fprintf(w, "%s\n", humanize.Commaf(float64(sum)))
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