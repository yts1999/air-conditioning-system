package service

import (
	"centralac/model"
	"centralac/serializer"
	"time"
)

// RoomShowService 获取房间信息的服务
type RoomShowService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Show 获取房间信息函数
func (service *RoomShowService) Show() serializer.Response {
	centerStatusLock.Lock()
	windSupplyLock.Lock()
	var room model.Room
	if err := model.DB.First(&room, service.RoomID).Error; err != nil {
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return serializer.ParamErr("房间信息不存在", err)
	}

	if room.WindSupply {
		var record model.Record
		if err := model.DB.First(&record, room.CurrentRecord).Error; err != nil {
			windSupplyLock.Unlock()
			centerStatusLock.Unlock()
			return serializer.SystemErr("无法查询当前记录", err)
		}
		minDur := float32(time.Now().Sub(record.StartTime).Minutes())
		var energy float32
		switch room.WindSpeed {
		case model.High:
			energy = minDur * 1.2
		case model.Medium:
			energy = minDur
		case model.Low:
			energy = minDur * 0.8
		}
		room.Energy += energy
		room.Bill += energy * 5.0
	}
	windSupplyLock.Unlock()
	centerStatusLock.Unlock()
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "获取房间信息成功"
	return resp
}
