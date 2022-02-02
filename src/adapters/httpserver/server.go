package httpserver

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	add "github.com/mikejeuga/go-exercises/src/domain"
	"log"
	"net/http"
	"strconv"
	"time"
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
	ctx := r.Context()

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintf(w, "Hello, %s!", "Go-Exercises")
		defer log.Printf("The server is healthy")
	case <-	ctx.Done():
		log.Print(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}

}

type AdderFuncServer func(num...int) int


func (f AdderFuncServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	 deadlineCtx, cancel := context.WithDeadline(r.Context(), time.Now().Add(6*time.Second))
	 defer cancel()
	 strings := r.URL.Query()["num"]
	 printResult(deadlineCtx, f(stringsToInts(strings)...), w)
}

func printResult(ctx context.Context, total int, w http.ResponseWriter) {
	deadline, _ := ctx.Deadline()

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "Total: %d", total)
		defer log.Printf("Added all the numbers before the %v:%v:%v deadline!", deadline.Hour(), deadline.Minute(), deadline.Second())
	case <-ctx.Done():
		log.Print(ctx.Err())
		defer log.Printf("Did nothing, the server has shut down")
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
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