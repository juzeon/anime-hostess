package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/reqstruct"
	"github.com/juzeon/anime-hostess/store"
)

func UserGenerate() include.Result {
	uid, token := store.CreateUser()
	return include.NewSuccessResult(gin.H{
		"uid":   uid,
		"token": token,
	})
}
func UserGetProgress(userID int, request reqstruct.HashUriRequest) include.Result {
	t, _ := store.GetUserFieldByID(userID, "progress:"+request.Hash).Float64()
	return include.NewSuccessResult(t)
}
func UserSetProgress(userID int, request reqstruct.UserSetProgressRequest) include.Result {
	store.SetUserFieldByID(userID, "progress:"+request.Hash, fmt.Sprint(request.Time))
	return include.NewOKResult()
}
func UserGetSearchText(userID int, request reqstruct.HashUriRequest) include.Result {
	text := store.GetUserFieldByID(userID, "searchText:"+request.Hash).String()
	if text == "" {
		_, _ = include.GetAllSeries(false)
		series, ok := include.Hash2SeriesMap[request.Hash]
		if !ok {
			return include.NewErrorResult("hash不存在")
		}
		text = series.Name
	}
	return include.NewSuccessResult(text)
}
func UserSetSearchText(userID int, request reqstruct.UserSetSearchTextRequest) include.Result {
	store.SetUserFieldByID(userID, "searchText:"+request.Hash, request.SearchText)
	return include.NewOKResult()
}
