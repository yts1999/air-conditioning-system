package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomTargetTempUpdateService 更改目标温度的服务
type RoomTargetTempUpdateService struct {
	RoomID     string  `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	TargetTemp float32 `form:"target_temp" json:"target_temp"`
}

// Update 更改目标温度函数
func (service *RoomTargetTempUpdateService) Update() serializer.Response {
	//检查房间号是否已经存在
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}
	centerStatusLock.RLock()
	if service.TargetTemp > highestTemp || service.TargetTemp < lowestTemp {
		centerStatusLock.RUnlock()
		return serializer.SystemErr("温度超出范围", nil)
	}
	centerStatusLock.RUnlock()
	// 更改目标温度
	if err := model.DB.Model(&room).Update("target_temp", service.TargetTemp).Error; err != nil {
		return serializer.DBErr("目标温度更改失败", err)
	}
	room.TargetTemp = service.TargetTemp
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "目标温度更改成功"
	return resp
}
