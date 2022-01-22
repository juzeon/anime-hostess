package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/router/reqstruct"
	"github.com/juzeon/anime-hostess/store"
)

func UserGenerate() include.Result {
	uid, token := store.CreateUser()
	return include.NewSuccessResult(gin.H{
		"uid":   uid,
		"token": token,
	})
}
func UserGetProgress(userID int, request reqstruct.UserGetProgressRequest) include.Result {
	t, _ := store.GetUserFieldByID(userID, "progress:"+request.Hash).Int()
	return include.NewSuccessResult(t)
}
func UserSetProgress(userID int, request reqstruct.UserSetProgressRequest) include.Result {
	store.SetUserFieldByID(userID, "progress:"+request.Hash, fmt.Sprint(request.Time))
	return include.NewOKResult()
}
