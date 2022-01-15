package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/router"
	"strconv"
)

func main() {
	mode := "release"
	if include.Config.Debug {
		mode = "debug"
	}
	gin.SetMode(mode)
	server := gin.Default()
	videoRouter := server.Group("/video")
	router.RegisterVideoRouter(videoRouter)
	err := server.Run(":" + strconv.Itoa(include.Config.Port))
	if err != nil {
		panic(err)
	}
}
