package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTaskList(ginCtx *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = int64(claim.Id)

	taskResp, err := taskService.GetTaskList(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"task" : taskResp.TaskList,
			"count" :taskResp.Count,
		},
	})
}

func CreateTask(ginCtx *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = int64(claim.Id)

	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": taskResp.TaskDetail})
}

func GetTaskDetail(ginCtx *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = int64(claim.Id)

	id, _ := strconv.Atoi(ginCtx.Param("id")) //获取task_id
	taskReq.Id = int64(id)

	taskResp, err := taskService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": taskResp.TaskDetail})
}

func UpdateTask(ginCtx *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = int64(claim.Id)

	id, _ := strconv.Atoi(ginCtx.Param("id")) //获取task_id
	taskReq.Id = int64(id)

	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": taskResp.TaskDetail})
}

func DeleteTask(ginCtx *gin.Context)  {
	var taskReq services.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))

	taskService := ginCtx.Keys["taskService"].(services.TaskService)
	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = int64(claim.Id)

	id, _ := strconv.Atoi(ginCtx.Param("id")) //获取task_id
	taskReq.Id = int64(id)

	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": taskResp.TaskDetail})
}