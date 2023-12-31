package app

import (
	"log"

	"electronic_diary/internal/domain/parent"
	ParentUC "electronic_diary/internal/domain/parent/usecase"
	"electronic_diary/internal/domain/student"
	StudentUC "electronic_diary/internal/domain/student/usecase"
	"electronic_diary/internal/domain/subject"
	SubjectUC "electronic_diary/internal/domain/subject/usecase"
	"electronic_diary/internal/domain/teacher"
	TeacherUC "electronic_diary/internal/domain/teacher/usecase"
	"electronic_diary/internal/domain/user"
	UserUC "electronic_diary/internal/domain/user/usecase"
	"electronic_diary/internal/services/auth"
	postgesql "electronic_diary/pkg/client/postgres"
	"electronic_diary/pkg/mailer"

	"electronic_diary/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	cfg      *config.Config
	pgClient *gorm.DB
	router   *gin.Engine

	userUC    user.UseCase
	subjectUC subject.UseCase
	parentUC  parent.UseCase
	studentUC student.UseCase
	teacherUC teacher.UseCase

	authService auth.Service
	mail        *mailer.Mailer
}

func NewApp(cfg *config.Config) *App {
	// Configs
	pgConfig := postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database)

	// Clients
	pgClient := postgesql.NewClient(pgConfig)
	mail, err := mailer.NewMailer(mailer.Config{
		From:     cfg.Mail.From,
		HOST:     cfg.Mail.Host,
		Port:     cfg.Mail.Port,
		Username: cfg.Mail.Username,
		Password: cfg.Mail.Password,
		SSL:      cfg.Mail.SSL,
	})
	if err != nil {
		log.Println("mail: ", err)
	}

	// UseCases
	userUC := UserUC.New(pgClient)
	subjectUC := SubjectUC.New(pgClient)
	parentUC := ParentUC.New(pgClient, userUC)
	studentUC := StudentUC.New(pgClient, userUC)
	teacherUC := TeacherUC.New(pgClient, userUC)

	// Services
	authService := auth.New(userUC, cfg.App, pgClient)

	router := gin.New()

	return &App{
		cfg:      cfg,
		pgClient: pgClient,
		router:   router,

		userUC:    userUC,
		subjectUC: subjectUC,
		parentUC:  parentUC,
		studentUC: studentUC,
		teacherUC: teacherUC,

		authService: authService,
		mail:        mail,
	}
}

func (a *App) Run() {
	a.setupHTTP()
}
