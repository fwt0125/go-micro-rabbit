package core

import (
	"task/model"
	"task/services"
)

func buildTask(item model.Task) *services.TaskModel {
	taskModel := services.TaskModel{
		Id:         int64(item.ID),
		Uid:        int64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status: int32(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}

