package main

import (
	"github.com/mrluizandre/gopportunities/config"
	"github.com/mrluizandre/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// initialize router
	router.Initialize()
}
