package store

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/juzeon/anime-hostess/include"
	"strconv"
)

var RDClient *redis.Client
var RDCtx context.Context
var RDPrefix = "ah_"

func ConnectRedisServer() {
	RDCtx = context.Background()
	RDClient = redis.NewClient(&redis.Options{
		Addr:     include.Config.RedisServer,
		Password: include.Config.RedisPassword,
	})
}

type RDField struct {
	data string
}

func (u RDField) MustInt() int {
	res, err := strconv.Atoi(u.data)
	if err != nil {
		panic(err)
	}
	return res
}
func (u RDField) Int() (int, error) {
	res, err := strconv.Atoi(u.data)
	if err != nil {
		return 0, err
	}
	return res, nil
}
func (u RDField) String() string {
	return u.data
}
