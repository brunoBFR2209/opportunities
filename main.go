package main

import (
	"api-opportunities/config"
	"api-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initializaton error: %v", err)
		return
	}

	router.Initialize()
}
