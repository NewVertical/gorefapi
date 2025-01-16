package core

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	Port int
}

func (s Server) StartServer(router *mux.Router) error {

	fmt.Println("Adding Handlers")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World"))
		if err != nil {
			return
		}
	})
	fmt.Printf("Starting server on port %d", s.Port)
	return http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.Port), router)
}
