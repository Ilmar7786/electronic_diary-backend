package http_delivery

import (
	adminSubjectsController "electronic_diary/internal/controller/http_delivery/admin_subjects__controller"
	adminUsersController "electronic_diary/internal/controller/http_delivery/admin_users__controller"
	"electronic_diary/internal/controller/http_delivery/auth_controller"
	subjectsController "electronic_diary/internal/controller/http_delivery/subjects__controller"
	userController "electronic_diary/internal/controller/http_delivery/user_controller"
	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/user"
	authService "electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

type Options struct {
	UserUC    user.UseCase
	SubjectUC subject.UseCase

	AuthService authService.Service
}

func Register(router *gin.RouterGroup, opt Options) {
	authController.Register(router, opt.AuthService)
	userController.Register(router, opt.AuthService, opt.UserUC)
	adminUsersController.Register(router, opt.AuthService, opt.UserUC)
	adminSubjectsController.Register(router, opt.AuthService, opt.SubjectUC)
	subjectsController.Register(router, opt.SubjectUC)
}
