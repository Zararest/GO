package sfss

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"gitlab.com/lgp/http-server-example/cmd/sfss/internal/config"
	_ "gitlab.com/lgp/http-server-example/docs"
	handler "gitlab.com/lgp/http-server-example/internal/handler/http"
	fileRepository "gitlab.com/lgp/http-server-example/internal/repository/file"
	"gitlab.com/lgp/http-server-example/internal/server"
	fileService "gitlab.com/lgp/http-server-example/internal/service/file"
)

// app is main service entry.
type app struct {
	// http server for listening and serving.
	server *http.Server
}

// NewApp create new app using config.Config. Init repos, service and server.
func NewApp(cfg config.Config) *app {
	log.Printf("create new write repository for files")
	writeRepository, err := fileRepository.NewWriteRepository(&cfg.FileRepository)
	if err != nil {
		panic(err)
	}
	log.Printf("create new read repository for files")
	readRepository, err := fileRepository.NewReadRepository(&cfg.FileRepository)
	if err != nil {
		panic(err)
	}

	log.Printf("create new service for files")
	service := fileService.NewFileService(readRepository, writeRepository)

	log.Printf("create new http server")
	r := mux.NewRouter()

	// Init swagger documentation handler.
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.Use(handler.LogMiddleware)
	handler := handler.NewHandler(r, service)
	server := server.NewServer(&cfg.Server, handler)

	return &app{
		server: server,
	}
}

// Run starts listen and server of http server.
func (a *app) Run() {
	log.Printf("run http server")
	if err := a.server.ListenAndServe(); err != nil {
		panic(err)
	}
}

// Close http connection.
func (a *app) Close() {
	log.Printf("stop http server")
	a.server.Close()
}
