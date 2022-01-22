package store

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/juzeon/anime-hostess/include"
	"strconv"
)

func CreateUser() (uid int, token string) {
	id := RDClient.Incr(RDCtx, RDPrefix+"id:user").Val()
	RDClient.SAdd(RDCtx, RDPrefix+"user", id)
	token = uuid.NewString()
	RDClient.HSet(RDCtx, RDPrefix+"user:"+strconv.Itoa(int(id)), "token", token)
	RDClient.HSet(RDCtx, RDPrefix+"token:user", token, id)
	return int(id), token
}
func Auth(ctx *gin.Context) {
	if id := GetUserIDByToken(ctx.GetHeader("Authorization")); id == 0 {
		ctx.JSON(200, include.NewErrorResult("token不存在或不正确"))
		ctx.Abort()
	} else {
		ctx.Set("UserID", id)
	}
}
func GetUserIDByToken(token string) int {
	idStr := RDClient.HGet(RDCtx, RDPrefix+"token:user", token).Val()
	id, _ := strconv.Atoi(idStr)
	return id
}
func GetUserFieldByID(id int, field string) RDField {
	return RDField{
		data: RDClient.HGet(RDCtx, RDPrefix+"user:"+strconv.Itoa(id), field).Val(),
	}
}
func SetUserFieldByID(id int, field string, value string) {
	RDClient.HSet(RDCtx, RDPrefix+"user:"+strconv.Itoa(id), field, value)
}
