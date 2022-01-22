package httpserver

import (
	"fmt"
	"github.com/gorilla/mux"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"github.com/mikejeuga/go-exercises/utils"
	"net/http"
)

type Server struct {
	math add.Adder
}

func NewServer(math add.Adder) *http.Server {
	s := Server{math: math}

	router := mux.NewRouter()
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
	total := f(utils.StringsToInts(strings)...)
	fmt.Fprintf(w, "Total: %d", total)
}

