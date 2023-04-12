package main

import (
	"electronic_diary/internal/app"
	"electronic_diary/internal/config"
)

func main() {
	cfg := config.GetConfig()

	App := app.NewApp(cfg)
	App.Run()
}
