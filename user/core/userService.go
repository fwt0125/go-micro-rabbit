package core

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"user/model"
	"user/services"
)

func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if model.DB.Where("user_name=?", req.UserName).First(&user).Error != nil {
		if gorm.ErrRecordNotFound == nil {
			resp.Code = 400
			return nil
		}
	}
	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		return errors.New("两次密码不一致")
	}
	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已存在")
	}
	user := model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassWord(req.Password); err != nil {
		return err
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func BuildUser(item model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: string(item.CreatedAt.Unix()),
		UpdatedAt: string(item.UpdatedAt.Unix()),
	}
	return &userModel
}
