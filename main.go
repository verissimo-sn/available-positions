package main

import (
	"github.com/verissimo-sn/available-positions/config"
	"github.com/verissimo-sn/available-positions/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("app")
	configError := config.Init()
	if configError != nil {
		logger.Errorf("Error initializing config: %v", configError)
		panic(configError)
	}
	router.Init()
}
