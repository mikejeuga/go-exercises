package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"github.com/mikejeuga/go-exercises/utils"
	"io"
	"net/http"
	"net/url"
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

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()


if r.Header.Get("Content-Type") == "application/json"{
	var m map[string][]string
	err = json.Unmarshal(bytes, &m)
	total := f(utils.StringsToInts(m["num"])...)
	fmt.Fprintf(w, "Total: %d", total)
	return
}

	if len(bytes) == 0 {
		strings := r.URL.Query()["num"]
		total := f(utils.StringsToInts(strings)...)
		fmt.Fprintf(w, "Total: %d", total)
		return
	}

	query, err := url.ParseQuery((string(bytes)))
	total := f(utils.StringsToInts(query["num"])...)
	fmt.Fprintf(w, "Total: %d", total)
}

