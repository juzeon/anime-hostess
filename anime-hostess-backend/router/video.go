package router

import (
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/service"
	"net/http"
)

func RegisterVideoRouter(rg *gin.RouterGroup) {
	rg.GET("/list", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, service.VideoList())
	})
	rg.GET("/stream/:hash", include.ValidateFields(func(valid *validation.Validation, ctx *gin.Context) {
		valid.Required(ctx.Param("hash"), "hash")
	}), func(ctx *gin.Context) {
		hash := ctx.Param("hash")
		service.VideoStream(hash, ctx)
	})
}
