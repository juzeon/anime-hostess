package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/router/reqstruct"
	"github.com/juzeon/anime-hostess/service"
	"github.com/juzeon/anime-hostess/store"
)

func RegisterUserRouters(user *gin.RouterGroup) {
	user.POST("/generate", func(ctx *gin.Context) {
		ctx.JSON(200, service.UserGenerate())
	})
	user.GET("/progress/:hash", store.Auth, func(ctx *gin.Context) {
		var req reqstruct.UserGetProgressRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		ctx.JSON(200, service.UserGetProgress(include.GetUserIDFromContext(ctx), req))
	})
	user.POST("/progress", store.Auth, func(ctx *gin.Context) {
		var req reqstruct.UserSetProgressRequest
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		ctx.JSON(200, service.UserSetProgress(include.GetUserIDFromContext(ctx), req))
	})
}
