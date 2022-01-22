package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/router/reqstruct"
	"github.com/juzeon/anime-hostess/service"
)

func RegisterBulletRouters(bullet *gin.RouterGroup) {
	bullet.GET("/search/:text", func(ctx *gin.Context) {
		var req reqstruct.BulletSearchRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		ctx.JSON(200, service.BulletSearch(req))
	})
	bullet.GET("/anime/:seasonID", func(ctx *gin.Context) {
		var req reqstruct.BulletAnimeRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		ctx.JSON(200, service.BulletAnime(req))
	})
	bullet.GET("/bullet/:cid", func(ctx *gin.Context) {
		var req reqstruct.BulletBulletRequest
		if err := ctx.ShouldBindUri(&req); err != nil {
			ctx.JSON(200, include.NewErrorResult(err.Error()))
			return
		}
		ctx.JSON(200, service.BulletBullet(req))
	})
}
