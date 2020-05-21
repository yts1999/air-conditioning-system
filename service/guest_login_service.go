package service

import (
	"centralac/model"
	"centralac/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GuestLoginService 管理用户登录的服务
type GuestLoginService struct {
	RoomID   string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	IDNumber string `form:"id_number" json:"id_number" binding:"required,min=18,max=18"`
}

// setSession 设置session
func (service *GuestLoginService) setSession(c *gin.Context, guest model.Guest) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("guest_id", guest.ID)
	s.Save()
}

// Login 用户登录函数
func (service *GuestLoginService) Login(c *gin.Context) serializer.Response {
	var guest model.Guest

	if err := model.DB.Where("room_id = ?", service.RoomID).First(&guest).Error; err != nil {
		return serializer.ParamErr("房间号或身份证号错误", nil)
	}

	if guest.CheckGuestIDNumber(service.IDNumber) == false {
		return serializer.ParamErr("房间号或身份证号错误", nil)
	}

	// 设置session
	service.setSession(c, guest)

	resp := serializer.BuildGuestResponse(guest)
	resp.Msg = "登录成功"
	return resp
}
