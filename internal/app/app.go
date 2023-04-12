package app

import (
	"electronic_diary/internal/domain/admin/repository"
	gorm_postgesql "electronic_diary/pkg/client/postgres"
	"fmt"
	"os"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type App struct {
	cfg *config.Config

	router *gin.Engine
}

func NewApp(cfg *config.Config) *App {
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
		os.Setenv(gin.EnvGinMode, gin.DebugMode)
	} else {

		gin.SetMode(gin.ReleaseMode)
		os.Setenv(gin.EnvGinMode, gin.ReleaseMode)
	}

	router := gin.Default()
	pgConfig := gorm_postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database)

	// Clients
	pgClient := gorm_postgesql.NewClient(pgConfig)

	// Modules
	//adminRepository := repository.NewAdminRepository(pgClient)
	pgClient.AutoMigrate(&repository.GormAdmin{})

	return &App{
		cfg:    cfg,
		router: router,
	}
}

func (a *App) Run() {
	//a.setupHTTP()
}

func (a *App) setupHTTP() {
	a.router.Use(cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	}))

	//prefix := a.router.Group("api")

	addr := fmt.Sprintf("%s:%s", a.cfg.HTTP.HOST, a.cfg.HTTP.PORT)
	if err := a.router.Run(addr); err != nil {
		return
	}
}
