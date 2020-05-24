package service

import (
	"centralac/model"
	"centralac/serializer"
)

// AdminRegisterService 管理管理员注册服务
type AdminRegisterService struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

// valid 验证表单
func (service *AdminRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.Admin{}).Where("username = ?", service.Username).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "该管理员已经注册",
		}
	}

	return nil
}

// Register 管理员注册
func (service *AdminRegisterService) Register() serializer.Response {
	admin := model.Admin{
		Username: service.Username,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := admin.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建管理员
	if err := model.DB.Create(&admin).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	resp := serializer.BuildAdminResponse(admin)
	resp.Msg = "注册成功"
	return resp
}
