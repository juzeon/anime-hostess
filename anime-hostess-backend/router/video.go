package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/service"
	"net/http"
)

func RegisterVideoRouter(rg *gin.RouterGroup) {
	rg.GET("/list", videoList)
}
func videoList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.VideoList())
}
