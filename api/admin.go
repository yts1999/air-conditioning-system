package api

import (
	"centralac/serializer"
	"centralac/service"
	"errors"

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

// LoginReq 登录请求
type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
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

// Login 登录接口
func Login(c *gin.Context) {
	var req LoginReq
	err := c.ShouldBind(req)
	if err != nil {
		c.JSON(200, ErrorResponse(err))
		return
	}
	if req.Role == "room" {
		service := service.GuestLoginService{
			RoomID:   req.Username,
			IDNumber: req.Password,
		}
		res := service.Login(c)
		c.JSON(200, res)
		return
	} else if req.Role == "admin" {
		service := service.AdminLoginService{
			Username: req.Username,
			Password: req.Password,
		}
		res := service.Login(c)
		c.JSON(200, res)
		return
	}
	c.JSON(200, ErrorResponse(errors.New("wrong role")))
	return

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
