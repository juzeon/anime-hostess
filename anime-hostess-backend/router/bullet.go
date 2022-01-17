package router

import (
	"github.com/beego/beego/v2/adapter/validation"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/service"
	"strconv"
)

func RegisterBulletRouters(bullet *gin.RouterGroup) {
	bullet.GET("/search/:text", include.ValidateFields(func(valid *validation.Validation, ctx *gin.Context) {
		valid.Required(ctx.Param("text"), "text")
	}), func(ctx *gin.Context) {
		ctx.JSON(200, service.BulletSearch(ctx.Param("text")))
	})
	bullet.GET("/anime/:seasonID", include.ValidateFields(func(valid *validation.Validation, ctx *gin.Context) {
		valid.Numeric(ctx.Param("seasonID"), "seasonID")
	}), func(ctx *gin.Context) {
		seasonID, _ := strconv.Atoi(ctx.Param("seasonID"))
		ctx.JSON(200, service.BulletAnime(seasonID))
	})
}
