package web

import (
	"canchitas-libres-field/internal/configuration"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	config  *configuration.Configuration
	handler *Handler
}

func NewServer(configuration *configuration.Configuration, handler *Handler) (*Server, error) {
	return &Server{
		config:  configuration,
		handler: handler,
	}, nil
}                                                                    

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler.ServeHTTP)
	fmt.Printf("Server listening on %s%s\n", s.config.SERVER.DOMAIN, s.config.SERVER.SERVER_PORT)
	err := http.ListenAndServe(s.config.SERVER.SERVER_PORT, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %s \n", err)
	}
}
