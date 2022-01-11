package add

import (
    "bufio"
    "fmt"
    "github.com/dustin/go-humanize"
    "io"
    "strconv"
)

func Sum(numbers ...int) int {
   total := 0
   for _, num := range numbers {
    total += num
   }
 return total
}

func Add(in io.Reader, out io.Writer) {
    var values []int

    scanner := bufio.NewScanner(in)
    scanner.Split(bufio.ScanLines)

    var numbers []string
    for scanner.Scan() {
        numbers = append(numbers, scanner.Text())
    }

    for _, val := range numbers {
        num, err := strconv.Atoi(val)
        if err != nil {
            num = 0
        }
        values = append(values, num)

    }

    sum := Sum(values...)
    fmt.Fprintf(out, "%s" , humanize.Commaf(float64(sum)))

}