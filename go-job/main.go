package main

import (
	"github.com/dev-araujo/go-job/config"
	r "github.com/dev-araujo/go-job/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialized error: %v", err)
		return

	}

	r.Initialize()

}
