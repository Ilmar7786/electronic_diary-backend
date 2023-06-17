package main

import (
	"electronic_diary/internal/app"
	"electronic_diary/internal/config"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	App := app.NewApp(cfg)
	App.Run()
}
