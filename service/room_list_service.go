package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomListService 获取所有房间信息的服务
type RoomListService struct {
}

// List 获取所有房间信息函数
func (service *RoomListService) List() serializer.Response {
	var rooms []model.Room
	if err := model.DB.Find(&rooms).Error; err != nil {
		return serializer.DBErr("获取所有房间信息失败", err)
	}
	resp := serializer.BuildRoomsResponse(rooms)
	resp.Msg = "获取所有房间信息成功"
	return resp
}
