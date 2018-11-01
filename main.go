package main

import (
	"github.com/health-insurance-api/app"
	"github.com/health-insurance-api/config"
)

func main() {
	config := config.GetConfig()

	app : = &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
