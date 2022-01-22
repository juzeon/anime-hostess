package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juzeon/anime-hostess/include"
	"github.com/juzeon/anime-hostess/mygrpc"
	"github.com/juzeon/anime-hostess/router"
	"github.com/juzeon/anime-hostess/store"
	"google.golang.org/grpc"
	"net"
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
	store.ConnectRedisServer()
	mode := "release"
	if include.Config.Debug {
		mode = "debug"
	}
	gin.SetMode(mode)
	server := gin.Default()
	server.Use(include.CORS())
	videoRouter := server.Group("/video")
	router.RegisterVideoRouters(videoRouter)
	bulletRouter := server.Group("/bullet")
	router.RegisterBulletRouters(bulletRouter)
	userRouter := server.Group("/user")
	router.RegisterUserRouters(userRouter)
	err := server.Run(include.Config.Listen)
	if err != nil {
		panic(err)
	}
}
func initializeGRPC() {
	listener, err := net.Listen("tcp", include.Config.Listen)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	mygrpc.RegisterBulletServer(grpcServer, mygrpc.BulletService{})
	fmt.Println("Started grpc server at " + include.Config.Listen)
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
