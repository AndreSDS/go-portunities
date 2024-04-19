package main

import (
	"github.com/andresds/go-portunities/internal/config"
	router "github.com/andresds/go-portunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// initialize logger
	logger = config.GetLogger("main")

	// initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err.Error())
		return
	}

	// initialize router
	router.InitRouter()

}
