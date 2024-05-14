package main

import (
	"phone-numbers-server/internal/config"
	"phone-numbers-server/internal/config/server"
	"phone-numbers-server/internal/logger"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cfg, err := config.GetParameters()
	handleError(err)

	log, err := logger.Create(cfg.Log)
	handleError(err)

	app, err := server.Create(cfg.Server, &log)
	handleError(err)

	app.Run()
	defer app.Close()
}