package server

import (
	"fmt"
	"net/http"

	"github.com/joaomarcuslf/qr-generator/handlers"
)

type Server struct {
	Port string
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (a *Server) Run() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/generator", handlers.GenerateQr)

	fmt.Println("Server running on port: " + a.Port)

	http.ListenAndServe(":"+a.Port, nil)
}
