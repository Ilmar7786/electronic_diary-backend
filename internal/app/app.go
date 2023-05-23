package app

import (
	"electronic_diary/internal/domain/auth"
	AuthModule "electronic_diary/internal/domain/auth/module"
	"electronic_diary/internal/domain/role"
	RoleModule "electronic_diary/internal/domain/role/module"
	"electronic_diary/internal/domain/user"
	UserModule "electronic_diary/internal/domain/user/module"

	postgesql "electronic_diary/pkg/client/postgres"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	cfg *config.Config

	router *gin.Engine

	roleModule role.Module
	userModule user.Module
	authModule auth.Module
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
	authModule := AuthModule.NewAuthModule(pgClient, userModule.GetUseCase())

	router := gin.New()

	return &App{
		cfg:    cfg,
		router: router,

		roleModule: roleModule,
		userModule: userModule,
		authModule: authModule,
	}
}

func (a *App) Run() {
	a.setupHTTP()
}
