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

	strings := r.URL.Query()["num"]
	query, err := url.ParseQuery((string(bytes)))
	if err != nil {
		panic(err)
	}

	if r.Header.Get("Content-Type") == "application/json"{
		f.jsonMath(w, bytes)
		return
	} else if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		f.formEncodeMath(w, strings)
		return
	} else {
		total := f(utils.StringsToInts(query["num"])...)
		fmt.Fprintf(w, "Total: %d", total)
	}

}

func (f AdderFuncServer) formEncodeMath(w http.ResponseWriter, strings []string) {
	total := f(utils.StringsToInts(strings)...)
	fmt.Fprintf(w, "Total: %d", total)
}
type MathPayload struct {
	Numbers []string `json:"nums"`
}

func (f AdderFuncServer) jsonMath(w http.ResponseWriter, bytes []byte) {
	var m MathPayload
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		panic(err)
	}
	total := f(utils.StringsToInts(m.Numbers)...)
	fmt.Fprintf(w, "Total: %d", total)
}

