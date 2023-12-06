package main

import (
	"log"

	"electronic_diary/internal/config"
	"electronic_diary/internal/constants"
	"electronic_diary/internal/domain/user/dto"
	UserUC "electronic_diary/internal/domain/user/usecase"
	postgesql "electronic_diary/pkg/client/postgres"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
)

func main() {
	cfg := config.GetConfig()

	pgConfig := postgesql.NewConfig(
		cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Database)

	// Clients
	pgClient := postgesql.NewClient(pgConfig)

	// UseCases
	userUC := UserUC.New(pgClient)

	_, err := userUC.Create(dto.CreateUserDTO{
		Surname:     readLine("Фамиля: "),
		Name:        readLine("Имя: "),
		Patronymic:  readLine("Отчество: "),
		Address:     readLine("Адрес: "),
		Phone:       readLine("Телефон: "),
		Email:       readLine("Почта: "),
		Password:    readLine("Пароль: ", input.WithEchoMode(input.EchoPassword)),
		Role:        constants.TeacherRole,
		IsSuperUser: true,
		IsActive:    true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("user successfully created")
}

func readLine(placeholder string, mode ...input.Option) string {
	val, err := prompt.New().Ask(placeholder).Input("", mode...)
	if err != nil {
		log.Fatalln(err)
	}

	return val
}
