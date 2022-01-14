package domain

type Adder interface {
   Add(numbers ...int) int
}

type AddProvider struct {}

func NewAddProvider() *AddProvider {
    return &AddProvider{}
}


func (a AddProvider) Add(numbers...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}