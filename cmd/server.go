package main

import (
	"k8s.io/client-go/kubernetes"
	"github.com/docker/docker/client"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"time"
)

type server struct {
	s       *http.Server
	m       *http.ServeMux
	kconfig *rest.Config
	kclient *kubernetes.Clientset
	dclient *client.Client
}

func newHTTPServer(mux *http.ServeMux) *http.Server {

	return &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

}

func newMux() *http.ServeMux {
	return &http.ServeMux{}
}

func NewServer() *server {
	m := newMux()

	kcfg := kubernetes_new_config("/home/quack/.kube/config")
	kcli := kubernetes_new_clientset(kcfg)
	dcli := docker_new_client()


	return &server{
		s:       newHTTPServer(m),
		m:       m,
		kclient: kcli,
		kconfig: kcfg,
		dclient: dcli,
	}
}

func (s *server) Serve() {
	log.Println("Listening And Serving...")
	s.Routes()
	ingress, err := s.kubernetes_new_ingress()

	if err != nil {
		s.LogError("Serve", err)
		return
	}
	s.LogMsg(ingress)

	err = s.s.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func (s *server) Routes() {
	s.m.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
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
