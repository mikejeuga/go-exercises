package domain

import (
    "io"
)

type Adder interface {
    Add(r io.Reader, w io.Writer)
}

type Service struct {}

func NewService() *Service {
    return &Service{}
}

func (s Service) Add(numbers...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}