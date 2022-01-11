package add

import (
    "bufio"
    "fmt"
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

func Add(in io.Reader, out io.Writer) {
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

    fmt.Fprintf(out, "%v" ,Sum(values...))

}