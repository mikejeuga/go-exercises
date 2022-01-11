package add

import (
    "bufio"
    "bytes"
    "fmt"
    "github.com/dustin/go-humanize"
    "io"
    "strconv"
    "strings"
)

func Sum(numbers ...int) int {
   total := 0
   for _, num := range numbers {
    total += num
   }
 return total
}

func Add(in []byte, out io.Writer) {
    values := GetData(in)
    toInts := StringsToInts(values)
    sum := Sum(toInts...)
    fmt.Fprintf(out, "%s\n", humanize.Commaf(float64(sum)))
}


func StringsToInts(numbers []string) []int {
    var values []int
    for _, val := range numbers {
        num, err := strconv.Atoi(val)
        if err != nil {
            num = 0
        }
        values = append(values, num)
    }
    return values
}

func GetData(in []byte) []string{
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