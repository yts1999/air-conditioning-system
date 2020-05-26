package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomShowService 获取房间信息的服务
type RoomShowService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Show 获取房间信息函数
func (service *RoomShowService) Show() serializer.Response {
	var room model.Room
	if err := model.DB.First(&room, service.RoomID).Error; err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "房间不存在",
			Error: err.Error(),
		}
	}
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "获取房间信息成功"
	return resp
}
