package service

import (
	"centralac/model"
	"centralac/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AdminLoginService 管理管理员登录的服务
type AdminLoginService struct {
	Username string `form:"username" json:"username" binding:"required,min=4,max=16"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=16"`
}

// setSession 设置session
func (service *AdminLoginService) setSession(c *gin.Context, admin model.Admin) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("admin_id", admin.ID)
	s.Save()
}

// Login 管理员登录函数
func (service *AdminLoginService) Login(c *gin.Context) serializer.Response {
	var admin model.Admin

	if err := model.DB.Where("username = ?", service.Username).First(&admin).Error; err != nil {
		return serializer.ParamErr("用户名错误", nil)
	}

	if admin.CheckPassword(service.Password) == false {
		return serializer.ParamErr("密码错误", nil)
	}

	// 设置session
	service.setSession(c, admin)

	resp := serializer.BuildUserResponse(admin)
	resp.Msg = "登录成功"
	return resp
}
