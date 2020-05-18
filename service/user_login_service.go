package service

import (
	"centralac/model"
	"centralac/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLoginService 管理用户登录的服务
type UserLoginService struct {
	RoomID   string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Password string `form:"password" json:"password" binding:"required,min=18,max=18"`
}

// setSession 设置session
func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

// Login 用户登录函数
func (service *UserLoginService) Login(c *gin.Context) serializer.Response {
	var user model.User

	if err := model.DB.Where("room_id = ?", service.RoomID).First(&user).Error; err != nil {
		return serializer.ParamErr("房间号或身份证号错误", nil)
	}

	if user.CheckPassword(service.Password) == false {
		return serializer.ParamErr("房间号或身份证号错误", nil)
	}

	// 设置session
	service.setSession(c, user)

	resp := serializer.BuildUserResponse(user)
	resp.Msg = "登录成功"
	return resp
}
