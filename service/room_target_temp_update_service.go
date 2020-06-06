package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomTargetTempUpdateService 更改目标温度的服务
type RoomTargetTempUpdateService struct {
	RoomID     string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	TargetTemp uint   `form:"target_temp" json:"target_temp"`
}

// Update 更改目标温度函数
func (service *RoomTargetTempUpdateService) Update() serializer.Response {
	//检查房间号是否已经存在
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}
	// 更改目标温度
	err := model.DB.Model(model.Room{}).
		Where("room_id = ?", service.RoomID).
		Update("target_temp", service.TargetTemp).Error
	if err != nil {
		return serializer.DBErr("目标温度更改失败", err)
	}
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "目标温度更改成功"
	return resp
}
