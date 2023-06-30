package http_delivery

import (
	"electronic_diary/internal/controller/http_delivery/admin__controller"
	"electronic_diary/internal/controller/http_delivery/auth__controller"
	parentController "electronic_diary/internal/controller/http_delivery/parent__controller"
	"electronic_diary/internal/controller/http_delivery/subjects__controller"
	teacherController "electronic_diary/internal/controller/http_delivery/teacher__controller"
	"electronic_diary/internal/domain/parent"
	"electronic_diary/internal/domain/student"
	"electronic_diary/internal/domain/subject"
	"electronic_diary/internal/domain/teacher"
	"electronic_diary/internal/domain/user"
	authService "electronic_diary/internal/services/auth"

	"github.com/gin-gonic/gin"
)

type Options struct {
	UserUC    user.UseCase
	SubjectUC subject.UseCase
	ParentUC  parent.UseCase
	TeacherUC teacher.UseCase
	StudentUC student.UseCase

	AuthService authService.Service
}

func Register(router *gin.RouterGroup, opt Options) {
	adminController.Register(router, opt.AuthService, opt.SubjectUC, opt.UserUC, opt.TeacherUC)
	authController.Register(router, opt.AuthService)
	subjectsController.Register(router, opt.SubjectUC)
	parentController.Register(router, opt.ParentUC)
	teacherController.Register(router, opt.TeacherUC)
}
