package add

import (
    "bufio"
    "fmt"
    "github.com/dustin/go-humanize"
    "io"
    "strconv"
    "strings"
)

func Add(numbers ...int) int {
   total := 0
   for _, num := range numbers {
    total += num
   }
 return total
}

func PrintAdd(in io.Reader, out io.Writer) {
    var values []int

    scanner := bufio.NewScanner(in)
    scanner.Scan()
    input := scanner.Text()

    split := strings.Split(input, " ")
    for _, val := range split {
        num, err := strconv.Atoi(val)
        if err != nil {
            num = 0
        }
        values = append(values, num)

    }

    sum := Add(values...)
    fmt.Fprintf(out, "%s" , humanize.Commaf(float64(sum)))

}