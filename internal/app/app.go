package app

import (
	"electronic_diary/internal/domain/user"
	UserUC "electronic_diary/internal/domain/user/usecase"
	"electronic_diary/internal/services/auth"
	postgesql "electronic_diary/pkg/client/postgres"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	cfg      *config.Config
	pgClient *gorm.DB
	router   *gin.Engine

	userUseCase user.UseCase

	authService auth.Service
}

func NewApp(cfg *config.Config) *App {
	// Configs
	pgConfig := postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database)

	// Clients
	pgClient := postgesql.NewClient(pgConfig)

	// UseCases
	userUC := UserUC.New(pgClient)

	// Services
	authService := auth.New(userUC, cfg.App)

	router := gin.New()

	return &App{
		cfg:      cfg,
		pgClient: pgClient,
		router:   router,

		userUseCase: userUC,

		authService: authService,
	}
}

func (a *App) Run() {
	a.setupHTTP()
}
