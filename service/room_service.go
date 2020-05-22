package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomService 管理房间的服务
type RoomService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Create 创建房间函数
func (service *RoomService) Create() serializer.Response {
	var room model.Room

	//检查房间号是否已经存在
	if !model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号已存在", nil)
	}

	room = model.Room{
		RoomID:     service.RoomID,
		SwitchTime: 0,
	}

	// 创建房间
	if err := model.DB.Create(&room).Error; err != nil {
		return serializer.ParamErr("房间创建失败", err)
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "房间创建成功"
	return resp
}

// Delete 删除房间函数
func (service *RoomService) Delete() serializer.Response {
	//检查房间号是否已经存在
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}

	//检查房间是否有客户
	var guest model.Guest
	if !model.DB.Where("room_id = ?", service.RoomID).First(&guest).RecordNotFound() {
		return serializer.ParamErr("房间中有客户", nil)
	}

	// 删除房间
	if err := model.DB.Where("room_id = ?", service.RoomID).Delete(&room).Error; err != nil {
		return serializer.ParamErr("房间删除失败", err)
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "房间删除成功"
	return resp
}
