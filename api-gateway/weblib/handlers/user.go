package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ginCtx gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(userReq))

	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
}

func UserLogin(ctx gin.Context) {
	var userReq services.UserRequest
	PanicIfUserError(ginCtx.Bind(userReq))

	userService := ginCtx.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(userResp.UserDetail.ID)
	ginCtx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "suss",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
