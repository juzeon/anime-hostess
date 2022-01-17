package router

import (
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/service"
	"net/http"
)

func RegisterVideoRouters(video *gin.RouterGroup) {
	video.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, service.VideoList())
	})
	video.GET("/stream/:hash", include.ValidateFields(func(valid *validation.Validation, ctx *gin.Context) {
		valid.Required(ctx.Param("hash"), "hash")
	}), func(ctx *gin.Context) {
		service.VideoStream(ctx.Param("hash"), ctx)
	})
}
