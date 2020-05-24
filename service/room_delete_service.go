package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomDeleteService 管理房间的服务
type RoomDeleteService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Delete 删除房间函数
func (service *RoomDeleteService) Delete() serializer.Response {
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
