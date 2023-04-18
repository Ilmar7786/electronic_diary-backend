package app

import (
	"fmt"
	"log"

	"electronic_diary/docs"
	"electronic_diary/internal/domain/admin"
	AdminModule "electronic_diary/internal/domain/admin/module"
	"electronic_diary/internal/domain/role"
	RoleModule "electronic_diary/internal/domain/role/module"
	"electronic_diary/internal/domain/user"
	UserModule "electronic_diary/internal/domain/user/module"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	postgesql "electronic_diary/pkg/client/postgres"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type App struct {
	cfg *config.Config

	router *gin.Engine

	roleModule  role.Module
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
	roleModule := RoleModule.NewRoleModule(pgClient)
	userModule := UserModule.NewUserModule(pgClient)
	adminModule := AdminModule.NewAdminModule(pgClient)

	router := gin.Default()

	return &App{
		cfg:    cfg,
		router: router,

		roleModule:  roleModule,
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

	if a.cfg.App.Debug {
		docs.SwaggerInfo.BasePath = fmt.Sprintf("/%s", apiPrefix)
		docs.SwaggerInfo.Host = addr
		a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

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
	a.roleModule.RegisterController(prefix)
	a.userModule.RegisterController(prefix)
	a.adminModule.RegisterController(prefix)

	if err := a.router.Run(addr); err != nil {
		log.Fatalln(err)
	}
}
