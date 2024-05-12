package main

import (
	"os"

	"gitlab.com/lgp/http-server-example/cmd/sfss/internal/app/sfss"
	"gitlab.com/lgp/http-server-example/cmd/sfss/internal/config"
)

// @title           Static File Storage Service (SFSS)
// @version         1.0
// @description     Service for store files in base dir

// @contact.email  andrianovartemii@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /
func main() {
	// Open config file.
	configFile, err := os.Open(*ConfigPathFlag)
	if err != nil {
		panic(err)
	}

	// Get app config from config file.
	cfg, err := config.Unmarshal(configFile)
	if err != nil {
		panic(err)
	}

	// Create new main app.
	app := sfss.NewApp(cfg)

	// Run and defer close main app.
	app.Run()
	defer app.Close()
}
