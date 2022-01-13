package main

import (
    "bufio"
    "bytes"
    "fmt"
    "github.com/dustin/go-humanize"
    "github.com/spf13/pflag"
    "io"
    "os"
    "strconv"
    "strings"
)


func main()  {
    var intCheck []int
    for _, val := range os.Args[1:] {
        num, _ := strconv.Atoi(val)
        intCheck = append(intCheck, num)
    }
    if len(os.Args[1:]) > 1 && Sum(intCheck...) > 0 {
        fmt.Fprintf(os.Stdout, "%s\n", humanize.Commaf(float64(Sum(intCheck...))))
        return
    }



    values:= []string{"input.csv"}
    stringP := pflag.StringSliceP("input-file", "i", values,"Enter the name of the files to be processed")
    pflag.Parse()


        for _, file := range *stringP {
            readFile, _ := os.ReadFile(file)
            defaultFile := bytes.NewReader(readFile)
            Add(defaultFile, os.Stdout)
        }
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