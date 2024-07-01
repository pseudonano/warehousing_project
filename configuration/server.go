package configuration

import (
	"log"
	"net/http"
	"time"
)

type Server interface {
	Run()
}

type server struct {
	*http.Server
}

func NewServer(addr string, handler http.Handler) Server {
	return &server{
		&http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
}

func (s *server) Run() {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		log.Printf("server running on: %v", s.Addr)
	}()
}
