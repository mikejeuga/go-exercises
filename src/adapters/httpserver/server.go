package httpserver

import (
	"fmt"
	"github.com/gorilla/mux"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"net/http"
	"strconv"
)

type Server struct {
	math add.Adder
}

func NewServer(math add.Adder) *http.Server {
	router := mux.NewRouter()
	s := Server{math: math}
	router.HandleFunc("/", s.home).Methods(http.MethodGet)
	router.Handle("/math", AdderFuncServer(s.math.Add)).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8099",
		Handler: router,
	}
}


func (s Server) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The math server is healthy")
}

type AdderFuncServer func(num...int) int


func (f AdderFuncServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	strings := r.URL.Query()["num"]
	total := f(stringsToInts(strings)...)
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