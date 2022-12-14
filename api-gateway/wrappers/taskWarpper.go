package wrappers

import (
	"api-gateway/services"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	"strconv"
)

func NewTask(id int64, name string) *services.TaskModel {
	return &services.TaskModel{
		Id: id,
		Title: name,
		Content: "响应越野",
		StartTime: 1000,
		EndTime: 1000,
		Status: 0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

// DefaultTask 降级处理
func DefaultTask(resp interface{})  {
	taskModels := make([]*services.TaskModel, 0)
	for i := 0; i < 10; i++ {
		taskModels = append(taskModels, NewTask(int64(i), "降级"+strconv.Itoa(20+int(i))))
	}
	result := resp.(*services.TaskResponse)
	result.TaskList = taskModels
}

type TaskWrapper struct {
	client.Client
}

func (wrapper *TaskWrapper) Call(ctx context.Context, req client.Request, resp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,   //熔断器请求阈值，默认20，意思是有20个请求才能进行错误百分比计算
		ErrorPercentThreshold:  50,   //错误百分比，当错误超过百分比时，直接进行降级处理，直至熔断器再次 开启，默认50%
		SleepWindow:            5000, //过多长时间，熔断器再次检测是否开启，单位毫秒ms（默认5秒）
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, resp)
	}, func(err error) error {
		return err
	})
}

// NewTaskWrapper  初始化Wrapper
func NewTaskWrapper(c client.Client) client.Client {
	return &TaskWrapper{c}
}