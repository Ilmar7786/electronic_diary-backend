package main

import (
	"flag"

	"electronic_diary/internal/app"
	"electronic_diary/internal/config"
)

var commandPathConfig string
var commandMigrate string

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	flag.StringVar(&commandPathConfig, "config", "configs/config.local.yaml", "path to config app")
	flag.StringVar(&commandMigrate, "migrate", "disable", "migration options: up, down, disable")

	flag.Parse()

	cfg := config.GetConfig(commandPathConfig)
	App := app.NewApp(cfg)

	if commandMigrate != "disable" {
		App.Migrate(commandMigrate)
		return
	}

	App.Run()
}
