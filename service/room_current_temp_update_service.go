package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomCurrentTempUpdateService 更新房间当前温度的服务
type RoomCurrentTempUpdateService struct {
	RoomID      string  `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	CurrentTemp float32 `form:"current_temp" json:"current_temp"`
}

// Update 更新房间当前函数
func (service *RoomCurrentTempUpdateService) Update() serializer.Response {
	//检查房间号是否已经存在
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}
	// 更新当前温度
	err := model.DB.Model(model.Room{}).
		Where("room_id = ?", service.RoomID).
		Update("current_temp", service.CurrentTemp).Error
	if err != nil {
		return serializer.DBErr("房间当前温度失败", err)
	}
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "房间当前温度更新成功"
	return resp
}
