package service

import (
	"centralac/model"
	"centralac/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	RoomID          string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Password        string `form:"password" json:"password" binding:"required,min=18,max=18"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=18,max=18"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的身份证号不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("room_id = ?", service.RoomID).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "房间已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	user := model.User{
		RoomID: service.RoomID,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	resp := serializer.BuildUserResponse(user)
	resp.Msg = "注册成功"
	return resp
}
