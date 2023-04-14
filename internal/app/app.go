package app

import (
	"fmt"
	"os"

	"electronic_diary/internal/domain/admin"
	AdminModule "electronic_diary/internal/domain/admin/module"
	"electronic_diary/internal/domain/user"
	UserModule "electronic_diary/internal/domain/user/module"

	postgesql "electronic_diary/pkg/client/postgres"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type App struct {
	cfg *config.Config

	router      *gin.Engine
	adminModule admin.Module
	userModule  user.Module
}

func NewApp(cfg *config.Config) *App {
	if cfg.App.Debug {
		gin.SetMode(gin.DebugMode)
		os.Setenv(gin.EnvGinMode, gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		os.Setenv(gin.EnvGinMode, gin.ReleaseMode)
	}

	// Configs
	pgConfig := postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database)

	// Clients
	pgClient := postgesql.NewClient(pgConfig)

	// Migrations
	runAutoMigrate(pgClient)

	// Modules
	adminModule := AdminModule.NewAdminModule(pgClient)
	userModule := UserModule.NewUserModule(pgClient)

	router := gin.Default()

	return &App{
		cfg:    cfg,
		router: router,

		adminModule: adminModule,
		userModule:  userModule,
	}
}

func (a *App) Run() {
	a.setupHTTP()
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

	prefix := a.router.Group("api")

	// Controllers
	a.userModule.RegisterController(prefix)
	a.adminModule.RegisterController(prefix)

	addr := fmt.Sprintf("%s:%s", a.cfg.HTTP.HOST, a.cfg.HTTP.PORT)
	if err := a.router.Run(addr); err != nil {
		return
	}
}
