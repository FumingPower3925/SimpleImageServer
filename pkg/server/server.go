package server

import (
	"log"
	"net/http"
	"time"

	"github.com/FumingPower3925/SimpleImageServer/pkg/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	server *http.Server
}

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	controller, err := controllers.New()
	if err != nil {
		return nil, err
	}
	r.Mount("/", controller)

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Close() error {
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
