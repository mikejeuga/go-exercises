package domain

import "context"

type Adder interface {
    Add(ctx context.Context, numbers ...int) int
}

type AdderFunc func(num...int) int

func (f AdderFunc) Add(ctx context.Context, numbers ...int) int {
    return f(numbers...)
}


var(
    Default = AdderFunc(Add)
)

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