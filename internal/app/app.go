package app

import (
	"electronic_diary/docs"
	"electronic_diary/internal/domain/admin"
	AdminModule "electronic_diary/internal/domain/admin/module"
	"electronic_diary/internal/domain/user"
	UserModule "electronic_diary/internal/domain/user/module"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"

	postgesql "electronic_diary/pkg/client/postgres"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type App struct {
	cfg *config.Config

	router      *gin.Engine
	userModule  user.Module
	adminModule admin.Module
}

func NewApp(cfg *config.Config) *App {
	// Configs
	pgConfig := postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database)

	// Clients
	pgClient := postgesql.NewClient(pgConfig)

	// Migrations
	runAutoMigrate(pgClient)

	// Modules
	userModule := UserModule.NewUserModule(pgClient)
	adminModule := AdminModule.NewAdminModule(pgClient)

	router := gin.Default()

	return &App{
		cfg:    cfg,
		router: router,

		userModule:  userModule,
		adminModule: adminModule,
	}
}

func (a *App) Run() {
	a.setupHTTP()
}

const apiPrefix = "api"

func (a *App) setupHTTP() {
	addr := fmt.Sprintf("%s:%s", a.cfg.HTTP.HOST, a.cfg.HTTP.PORT)

	docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s", apiPrefix)
	docs.SwaggerInfo.Host = addr
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	a.router.Use(cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	}))

	prefix := a.router.Group(apiPrefix)

	// Controllers
	a.userModule.RegisterController(prefix)
	a.adminModule.RegisterController(prefix)

	if err := a.router.Run(addr); err != nil {
		log.Fatalln(err)
	}
}
