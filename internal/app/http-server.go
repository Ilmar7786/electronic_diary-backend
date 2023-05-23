package app

import (
	"fmt"
	"log"

	"electronic_diary/docs"

	cors "github.com/rs/cors/wrapper/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *App) setupHTTP() {
	addr := fmt.Sprintf("%s:%s", a.cfg.HTTP.HOST, a.cfg.HTTP.PORT)
	prefix := a.router.Group(a.cfg.HTTP.PrefixAPI)

	if a.cfg.App.Debug {
		docs.SwaggerInfo.Title = a.cfg.Swagger.Title
		docs.SwaggerInfo.Version = a.cfg.Swagger.Version
		docs.SwaggerInfo.Host = addr
		docs.SwaggerInfo.BasePath = "/" + a.cfg.HTTP.PrefixAPI
		docs.SwaggerInfo.Schemes = a.cfg.Swagger.Schemes
		prefix.GET(fmt.Sprintf("/%s/*any", a.cfg.Swagger.Path), ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err := a.router.SetTrustedProxies(a.cfg.HTTP.Proxy)
	if err != nil {
		log.Fatalln(err)
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

	// Controllers
	a.roleModule.RegisterController(prefix)
	a.userModule.RegisterController(prefix)
	a.authModule.RegisterController(prefix)

	if err := a.router.Run(addr); err != nil {
		log.Fatalln(err)
	}
}
