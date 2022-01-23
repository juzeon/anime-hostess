package service

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/reqstruct"
)

func VideoList() include.Result {
	series, err := include.GetAllSeries(true)
	if err != nil {
		return include.NewErrorResult(err)
	}
	return include.NewSuccessResult(series)
}
func VideoDetail(request reqstruct.HashUriRequest) include.Result {
	_, _ = include.GetAllSeries(false)
	series, ok := include.Hash2SeriesMap[request.Hash]
	if !ok {
		return include.NewErrorResult("hash不存在")
	}
	return include.NewSuccessResult(series)
}
func VideoStream(request reqstruct.HashUriRequest, ctx *gin.Context) {
	_, _ = include.GetAllSeries(false)
	path, ok := include.Hash2PathMap[request.Hash]
	if !ok {
		ctx.JSON(200, include.NewErrorResult("hash不存在"))
		return
	}
	ctx.File(path)
}
