package api

import (
	"centralac/serializer"
	"centralac/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AdminRegister 管理员注册接口
func AdminRegister(c *gin.Context) {
	var service service.AdminRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminLogin 管理员登录接口
func AdminLogin(c *gin.Context) {
	var service service.AdminLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// AdminLogout 管理员登出
func AdminLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
