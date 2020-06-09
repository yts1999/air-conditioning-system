package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomWindSpeedUpdateService 更改风速的服务
type RoomWindSpeedUpdateService struct {
	RoomID    string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	WindSpeed uint   `form:"wind_speed" json:"wind_speed"`
}

// Update 更改风速函数
func (service *RoomWindSpeedUpdateService) Update() serializer.Response {
	//检查房间号是否已经存在
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}

	// 当前正在送风，则结束本次送风请求，重新开始送风
	if room.WindSupply {
		resp := stopWindSupply(&room)
		if resp.Code != 0 {
			return resp
		}

		// 更改风速
		if err := model.DB.Model(&room).Update("wind_speed", service.WindSpeed).Error; err != nil {
			return serializer.DBErr("风速更改失败", err)
		}
		room.WindSpeed = service.WindSpeed

		resp = windSupply(&room)
		if resp.Code != 0 {
			return resp
		}
	} else {
		// 更改风速
		if err := model.DB.Model(&room).Update("wind_speed", service.WindSpeed).Error; err != nil {
			return serializer.DBErr("风速更改失败", err)
		}
		room.WindSpeed = service.WindSpeed
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "风速更改成功"
	return resp
}
