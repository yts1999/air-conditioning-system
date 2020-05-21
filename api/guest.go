package api

import (
	"centralac/model"
	"centralac/serializer"
	"centralac/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GuestRegister 房客注册接口
func GuestRegister(c *gin.Context) {
	var service service.GuestRegisterService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// GuestLogin 房客登录接口
func GuestLogin(c *gin.Context) {
	var service service.GuestLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CurrentGuest 获取当前房客
func CurrentGuest(c *gin.Context) *model.Guest {
	if guest, _ := c.Get("guest"); guest != nil {
		if u, ok := guest.(*model.Guest); ok {
			return u
		}
	}
	return nil
}

// GuestMe 房客详情
func GuestMe(c *gin.Context) {
	guest := CurrentGuest(c)
	res := serializer.BuildGuestResponse(*guest)
	c.JSON(200, res)
}

// GuestLogout 房客登出
func GuestLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}
