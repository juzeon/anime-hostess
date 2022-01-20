package service

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
)

func VideoList() include.Result {
	series, err := include.GetAllSeries(false)
	if err != nil {
		return include.NewErrorResult(err)
	}
	return include.NewSuccessResult(series)
}
func VideoDetail(hash string) include.Result {
	series, ok := include.Hash2SeriesMap[hash]
	if !ok {
		_, _ = include.GetAllSeries(true)
		series, ok = include.Hash2SeriesMap[hash]
	}
	if !ok {
		return include.NewErrorResult("hash不存在")
	}
	return include.NewSuccessResult(series)
}
func VideoStream(hash string, ctx *gin.Context) {
	path, ok := include.Hash2PathMap[hash]
	if !ok {
		_, _ = include.GetAllSeries(true)
		path, ok = include.Hash2PathMap[hash]
	}
	if !ok {
		ctx.JSON(200, include.NewErrorResult("hash不存在"))
		return
	}
	ctx.File(path)
}
