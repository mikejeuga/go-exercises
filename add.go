package main

import (
    "bufio"
    "bytes"
    "fmt"
    "github.com/dustin/go-humanize"
    "io"
    "os"
    "strconv"
    "strings"
)


func main()  {

    defaultBytes, _ := os.ReadFile("input.csv")
    defaultReader := bytes.NewReader(defaultBytes)

    //if there is no arguments given
    noArgs := len(os.Args) < 2
    if noArgs {
        Add(defaultReader, os.Stdout)
        return
    }

    //if a file name is given
    inputTxt, _ := os.ReadFile(os.Args[1])
    fileNameOnly := len(inputTxt) > 1 && len(os.Args) < 3
    if fileNameOnly {
        Add(bytes.NewReader(inputTxt), os.Stdout)
        return
    }

    //if numbers are given to the command-line
    ints := StringsToInts(os.Args[1:])
    Sum(ints...)
    fmt.Fprintf(os.Stdout, "%s\n", humanize.Commaf(float64(Sum(ints...))))

}


func Sum(numbers ...int) int {
   total := 0
   for _, num := range numbers {
    total += num
   }
 return total
}

func Add(in io.Reader, out io.Writer) {
    readAll, err := io.ReadAll(in)
    if err != nil {
        panic("it doesn't read io.ReadAll")
    }

    values := getData(readAll)
    toInts := StringsToInts(values)
    sum := Sum(toInts...)
    fmt.Fprintf(out, "%s\n", humanize.Commaf(float64(sum)))
}


func StringsToInts(numbers []string) []int {
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