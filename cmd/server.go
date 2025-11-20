package main

import (
	"net/http"
	"log"
	"time"
)

type server struct {
	s *http.Server
	m *http.ServeMux
}


func newHTTPServer(mux *http.ServeMux) *http.Server {

	return &http.Server{
		Addr: ":8080",
		Handler: mux,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

}

func newMux() *http.ServeMux {
	return &http.ServeMux{}
}


func NewServer() *server {
	m := newMux()
	return &server{
		s: newHTTPServer(m),
		m: m,
	}
}


func (s *server) Serve() {
	log.Println("Listening And Serving...")
	s.Routes()
	err := s.s.ListenAndServe()

	if err != nil {
		panic(err)
	}
}


func (s *server) Routes() {
	s.m.HandleFunc("GET /health", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
}
