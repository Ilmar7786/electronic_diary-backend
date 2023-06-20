package http_delivery

import (
	adminController "electronic_diary/internal/controller/http_delivery/admin_controller"
	"electronic_diary/internal/controller/http_delivery/auth_controller"
	subjectController "electronic_diary/internal/controller/http_delivery/subjects"
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
	adminController.Register(router, opt.AuthService, opt.UserUC)
	subjectController.Register(router, opt.AuthService, opt.SubjectUC)
}
