package httpserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Server struct {

}

func (s Server) Add(w http.ResponseWriter, r *http.Request) {

}

func NewServer() *Server {
	return &Server{}
}

type Adder interface {
	Add(ctx context.Context, numbers ...int) int
}

type AdderFunc func(num...int) int

func (f AdderFunc) Add(ctx context.Context, numbers ...int) int {
	return f(numbers...)
}



func (s Server) Run() error {
	router := mux.NewRouter()
	router.HandleFunc("/", s.Home).Methods(http.MethodGet)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("THE MATH SERVER IS RUNNING!!!")

	return server.ListenAndServe()

}


func (s Server) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The math server is healthy")
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