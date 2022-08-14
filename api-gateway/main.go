package main

import (
	"api-gateway/config"
	"api-gateway/services"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
	"time"
)

func main() {
	config.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress),
	)
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	userService := services.NewUserService("rpcUserService", userMicroService.Client())

	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)
	taskService := services.NewTaskService("rpcTaskService", taskMicroService.Client())

	server := web.NewService(
		web.Name("httpServer"),
		web.Address(":3000"),
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	server.Init()
	server.Run()

}
