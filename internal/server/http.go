package server

import (
	"context"
	"log"
	"net/http"
	"time"
)


type Server struct {
	httpServer *http.Server
}

func New(addr string, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: addr,
			Handler: handler,
			ReadTimeout: 10*time.Second,
			WriteTimeout: 10* time.Second,
			IdleTimeout: 60*time.Second,
		},
	}
}

func(s *Server) Start(){
	log.Println("HTTP server started on", s.httpServer.Addr)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed{
		log.Fatal(err)
	}
}


func(s *Server) Shutdown(ctx context.Context)error{
	log.Println("Shutting down HTTP server...")
	return s.httpServer.Shutdown(ctx)
}