package domain

import (
    "strconv"
)

var(
    Default = AdderFunc(Add)
)

type Adder interface {
    Add(numbers ...int) int
}

type AdderFunc func(num...int) int

func (f AdderFunc) Add(numbers ...int) int {
    return f(numbers...)
}


func Add(numbers...int) int {
    total := 0
    for _, num := range Uniq(numbers) {
        total += num
    }

    return total
}

func Uniq(in []int) []int {
    m := make(map[int]bool)
    uniqInts := []int{}
    for _, num := range in {
        if _, val := m[num]; !val {
            m[num] = true
            uniqInts = append(uniqInts, num)
        }
    }
    return uniqInts
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