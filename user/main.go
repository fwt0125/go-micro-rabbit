package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"user/config"
	"user/core"
	"user/services"
)

func main() {
	config.Init()
	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress),
	)
	//微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.EtcdMicroAddress),
		micro.Registry(etcdReg),
	)

	//初始化
	microService.Init()
	//服务注册
	err := services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	if err != nil {
		panic(err)
	}
	_ = microService.Run()
}
