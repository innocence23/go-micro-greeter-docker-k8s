package main

import (
	"context"
	greeter "greeter-client/proto"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

var client greeter.GreeterService

func main() {

	service := micro.NewService(
		micro.Name("go.micro.cli.greeter"),
		// 服务版本
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("etcd:2379"),
		)), //etcd注册
	)

	service.Init()

	client = greeter.NewGreeterService("go.micro.srv.greeter", service.Client())

	route := gin.Default()

	// When GET /greeter, handle the request with the Greet function we create above
	route.GET("/greeter", Greet)
	if err := route.Run("0.0.0.0:3000"); err != nil {
		log.Print(err.Error())
	}
}

func Greet(ctx *gin.Context) {
	name := ctx.Query("name") // ctx.Query will return the GET request query string
	log.Println("Client handle Greet, name =", name)

	// Client request the RPC server for response
	res, err := client.Greet(context.TODO(), &greeter.Request{Name: name})
	if err != nil {
		log.Print(err.Error())
		// return with 400 error code and error message
		ctx.JSON(400, gin.H{"message": "server error"})
		return
	}

	// return 200 success code and the response from server
	ctx.JSON(200, gin.H{"message": res.Greeting})
}
