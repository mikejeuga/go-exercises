package httpserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Server struct {
	math Adder
}

func NewServer(math Adder) *http.Server {
	router := mux.NewRouter()
	s := Server{math: math}
	router.HandleFunc("/", s.home).Methods(http.MethodGet)
	router.HandleFunc("/math", s.add).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8099",
		Handler: router,
	}
}

type Adder interface {
	Add(ctx context.Context, numbers ...int) int
}

type AdderFunc func(num...int) int


func (f AdderFunc) Add(ctx context.Context, numbers ...int) int {
	return f(numbers...)
}



func (s Server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The math server is healthy")
}


func (s Server) add(w http.ResponseWriter, r *http.Request) {
	strings := r.URL.Query()["num"]
	total := s.math.Add(r.Context(), stringsToInts(strings)...)
	fmt.Fprintf(w, "Total: %d", total)
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