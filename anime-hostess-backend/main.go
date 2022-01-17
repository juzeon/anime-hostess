package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/mygrpc"
	"github.com/juzeon/anime-hostess/router"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

func main() {
	if include.Config.AsGRPC {
		initializeGRPC()
	} else {
		initializeMaster()
	}
}
func initializeMaster() {
	mygrpc.ConnectServer()
	mode := "release"
	if include.Config.Debug {
		mode = "debug"
	}
	gin.SetMode(mode)
	server := gin.Default()
	videoRouter := server.Group("/video")
	router.RegisterVideoRouters(videoRouter)
	bulletRouter := server.Group("/bullet")
	router.RegisterBulletRouters(bulletRouter)
	err := server.Run(":" + strconv.Itoa(include.Config.Port))
	if err != nil {
		panic(err)
	}
}
func initializeGRPC() {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(include.Config.Port))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	mygrpc.RegisterBulletServer(grpcServer, mygrpc.BulletService{})
	fmt.Println("Started grpc server at port " + strconv.Itoa(include.Config.Port))
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
