package server

import (
	"phone-numbers-server/internal/config"
	"phone-numbers-server/internal/logger"
)

type Server struct {
	// FIXME
}

func (*Server) Run() {

}

func (*Server) Close() {

}

func Create(cfg config.ServerConfig, log *logger.Logger) (*Server, error) {
	return nil, nil //FIXME
}