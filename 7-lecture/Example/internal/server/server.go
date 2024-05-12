package server

import "net/http"

// NewServer returns http.Server by Config and handlers.
func NewServer(cfg *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    cfg.Host,
		Handler: handler,
	}
}
