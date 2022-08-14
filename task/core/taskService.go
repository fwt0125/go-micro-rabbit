package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/streadway/amqp"
	"task/model"
	"task/services"
)

// CreateTask 将消息放入队列
func (*TaskService)CreateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	channel, err := model.MQ.Channel()
	if err != nil {
			errors.New("MQ error" + err.Error())
	}
	declare, _ := channel.QueueDeclare("task_queue", true, false, false, false, nil)
	marshal, err := json.Marshal(req)
	err = channel.Publish("", declare.Name, false, false, amqp.Publishing{
		DeliveryMode:    amqp.Persistent,
		ContentEncoding: "application/json",
		Body:            marshal,
	})
	if err != nil {
		errors.New("MQ Publish" + err.Error())
	}
	return nil
}

func (*TaskService)GetTaskList(ctx context.Context, req *services.TaskRequest, resp *services.TaskResponse) error  {
	if req.Limit == 0 {
		req.Limit = 10
	}

	var taskData  []model.Task
	var count int64
	err := model.DB.Offset(int(req.Start)).Limit(int(req.Limit)).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return err
	}

	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)
	var taskRes []*services.TaskModel
	for _, datum := range taskData {
		taskRes = append(taskRes, buildTask(datum))
	}
	resp.TaskList = taskRes
	resp.Count = count
	return nil
}

func (*TaskService)GetTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := buildTask(taskData)
	resp.TaskDetail = taskRes
	return nil
}


func (*TaskService)UpdateTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.Model(model.Task{}).Where("id=? and uid=?",req.Id,req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content
	model.DB.Save(&taskData)
	resp.TaskDetail = buildTask(taskData)
	return nil
}

func (*TaskService)DeleteTask(ctx context.Context, req *services.TaskRequest, resp *services.TaskDetailResponse) error {
	model.DB.Model(model.Task{}).Where("id=? and uid=?",req.Id,req.Uid).Delete(&model.Task{})
	return nil
}