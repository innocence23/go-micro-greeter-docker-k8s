package main

import (
	"context"
	greeter "greeter-service/proto"
	"log"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

type GreeterSerivce struct{}

func (g *GreeterSerivce) Greet(ctx context.Context, req *greeter.Request, resp *greeter.Response) error {
	log.Println("Greeter service handle Greet", req.Name)
	resp.Greeting = "Hello, " + req.Name
	return nil
}

func main() {

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		// 服务版本
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs("etcd:2379"),
		)), //etcd注册
	)

	service.Init()

	greeter.RegisterGreeterHandler(service.Server(), new(GreeterSerivce))

	if err := service.Run(); err != nil {
		log.Fatal("run error")
	}

}
