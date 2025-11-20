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
		s.JSON(w, map[string]string{"status": "ok"}, 200)
	})

	s.m.HandleFunc("POST /v1/deploy", s.handlePostDeploy)
}


func (s *server) LogError(f string, err error) {
	log.Printf("ERROR::%v::ERROR\n%v\n", f, err.Error())
}

func (s *server) LogMsg(msg any) {
	log.Println("MESSAGE::", msg)
}

func (s *server) LogRequest(r *http.Request) {
	log.Printf("%s   %s   %s   at   %s", r.Method, r.URL.Path, r.Host, time.Now().Format("2006-01-02 15:04:05"))
}
