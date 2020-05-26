package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomCreateService 创建房间的服务
type RoomCreateService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Create 创建房间函数
func (service *RoomCreateService) Create() serializer.Response {
	var room model.Room

	//检查房间号是否已经存在
	if !model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号已存在", nil)
	}

	room = model.Room{
		RoomID: service.RoomID,
	}

	// 创建房间
	if err := model.DB.Create(&room).Error; err != nil {
		return serializer.ParamErr("房间创建失败", err)
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "房间创建成功"
	return resp
}
