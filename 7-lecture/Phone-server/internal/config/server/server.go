package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"phone-numbers-server/internal/config"
	"phone-numbers-server/internal/logger"
)

type server struct {
	server *http.Server
	log    *logger.Logger
}

func (s *server) Run() {
	s.log.Print("server startup", 1)
	if err := s.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *server) Close() {
	s.log.Print("server stop", 1)
	s.server.Close()
}

func createServer(r *mux.Router, cfg config.ServerConfig, log *logger.Logger) (*server, error) {
	return &server{&http.Server{Addr: cfg.Host, Handler: r}, log}, nil
}

func Create(cfg config.ServerConfig, log *logger.Logger) (*server, error) {
	log.Print("creating a server", 1)
	log.Print("creating router", 2)
	r := mux.NewRouter()

	initHandlers(r)
	return createServer(r, cfg, log)
}