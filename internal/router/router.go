package router

import (
	"ccgin/configs"
	"ccgin/internal/codes"
	"ccgin/internal/handler/user"
	"ccgin/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.Engine) {
	prefixApi := g.Group("/api")

	prefixApi.GET("/ping", func(ctx *gin.Context) {
		codes.Success.EmptyResponse(ctx)
	})

	auther := middleware.Auth(configs.DB(), configs.Rds(), configs.GetConfig().JWT.Secret)

	userHandler := user.New()
	userGroup := prefixApi.Group("/user")
	{
		userGroup.POST("/login", userHandler.Login())
	}
	userAuthGroup := prefixApi.Group("/user", auther)
	{
		userAuthGroup.GET("/info", userHandler.Info())
	}

}
