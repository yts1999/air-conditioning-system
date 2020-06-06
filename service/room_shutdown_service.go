package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomShutdownService 从控机关机的服务
type RoomShutdownService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Shutdown 从控机关机函数
func (service *RoomShutdownService) Shutdown() serializer.Response {
	var room model.Room
	if err := model.DB.First(&room, service.RoomID).Error; err != nil {
		return serializer.Err(404, "房间信息不存在", err)
	}

	// 从控机关机
	err := model.DB.Model(model.Room{}).
		Where("room_id = ?", service.RoomID).
		Update("power_on", false).Error
	if err != nil {
		return serializer.DBErr("从控机关机失败", err)
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "从控机关机成功"
	return resp
}
