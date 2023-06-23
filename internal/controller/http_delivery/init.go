package http_delivery

import (
	"electronic_diary/internal/controller/http_delivery/admin__controller"
	"electronic_diary/internal/controller/http_delivery/auth__controller"
	"electronic_diary/internal/controller/http_delivery/subjects__controller"
	"electronic_diary/internal/controller/http_delivery/user__controller"
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
	adminController.Register(router, opt.AuthService, opt.SubjectUC, opt.UserUC)
	authController.Register(router, opt.AuthService)
	userController.Register(router, opt.AuthService, opt.UserUC)
	subjectsController.Register(router, opt.SubjectUC)
}
