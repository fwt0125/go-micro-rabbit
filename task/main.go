package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"task/config"
	"task/core"
	"task/services"
)

func main()  {

	config.Init()
	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.EtcdAddress),
	)
	//微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"),
		micro.Address(config.EtcdMicroAddress),
		micro.Registry(etcdReg),
	)

	//初始化
	microService.Init()
	//服务注册
	err := services.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	if err != nil {
		panic(err)
	}
	_ = microService.Run()
}
