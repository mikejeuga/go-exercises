package httpserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {}

func NewServer() *Server {
	return &Server{}
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

