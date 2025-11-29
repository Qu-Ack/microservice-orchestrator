package main

import (
	"database/sql"
	"github.com/docker/docker/client"
	_ "github.com/lib/pq"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"sync"
	"time"
)

type server struct {
	s        *http.Server
	m        *http.ServeMux
	kconfig  *rest.Config
	kclient  *kubernetes.Clientset
	dclient  *client.Client
	db       *sql.DB
	mu       sync.Mutex
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

func newDB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=kube_orch sslmode=disable"	
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic("couldn't reach postgres DB")
	}

	return db
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
		s:        newHTTPServer(m),
		m:        m,
		kclient:  kcli,
		kconfig:  kcfg,
		dclient:  dcli,
		db:       newDB(),
		mu:       sync.Mutex{},
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
	s.m.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		s.JSON(w, map[string]string{"status": "ok"}, 200)
	})

	s.m.HandleFunc("POST /v1/deploy", s.MiddlewareExtractCookie(s.handlePostDeploy))
	s.m.HandleFunc("GET /v1/deploy", s.MiddlewareExtractCookie(s.GetDeployments))
	s.m.HandleFunc("PUT /v1/deploy", s.handlePutDeploy)

	s.m.HandleFunc("POST /v1/user/register", s.registerUser)
	s.m.HandleFunc("POST /v1/user/login", s.LogUser)
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
