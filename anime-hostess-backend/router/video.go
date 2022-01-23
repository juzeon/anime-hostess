package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/reqstruct"
	"github.com/juzeon/anime-hostess/service"
)

func RegisterVideoRouters(video *gin.RouterGroup) {
	video.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(200, service.VideoList())
	})
	video.GET("/detail/:hash", func(ctx *gin.Context) {
		var req reqstruct.HashUriRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		ctx.JSON(200, service.VideoDetail(req))
	})
	video.GET("/stream/:hash", func(ctx *gin.Context) {
		var req reqstruct.HashUriRequest
		if err := ctx.BindUri(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		service.VideoStream(req, ctx)
	})
}
