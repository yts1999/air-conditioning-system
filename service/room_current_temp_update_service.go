package service

import (
	"centralac/model"
	"centralac/serializer"
	"time"
)

// RoomCurrentTempUpdateService 更新房间当前温度的服务
type RoomCurrentTempUpdateService struct {
	RoomID      string  `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	CurrentTemp float32 `form:"current_temp" json:"current_temp"`
}

// Update 更新房间当前温度函数
func (service *RoomCurrentTempUpdateService) Update() serializer.Response {
	//检查房间号是否已经存在
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}
	// 更新当前温度
	if err := model.DB.Model(&room).Update("current_temp", service.CurrentTemp).Error; err != nil {
		return serializer.DBErr("房间当前温度失败", err)
	}
	room.CurrentTemp = service.CurrentTemp
	centerStatusLock.Lock()
	windSupplyLock.Lock()
	if !centerPowerOn {
		if room.WindSupply {
			stopWindSupply(&room)
		}
		if waitStatus[room.RoomID] {
			for i := waitList.Front(); i != nil; i = i.Next() {
				if i.Value == room.RoomID {
					waitList.Remove(i)
					delete(waitStatus, room.RoomID)
					break
				}
			}
		}
	}
	windSupplyLock.Unlock()
	var StartTime time.Time
	if room.WindSupply {
		var record model.Record
		if err := model.DB.First(&record, room.CurrentRecord).Error; err != nil {
			centerStatusLock.Unlock()
			return serializer.SystemErr("无法查询当前记录", err)
		}
		StartTime = record.StartTime
		curTime := time.Now()
		minDur := float32(curTime.Sub(record.StartTime).Minutes())
		var energy float32
		switch room.WindSpeed {
		case model.High:
			energy = minDur * 1.2
		case model.Medium:
			energy = minDur
		case model.Low:
			energy = minDur * 0.8
		}
		recordNew := make(map[string]interface{})
		recordNew["end_time"] = curTime
		recordNew["end_temp"] = service.CurrentTemp
		recordNew["energy"] = energy
		recordNew["bill"] = energy * 5.0
		if err := model.DB.Model(&record).Updates(recordNew).Error; err != nil {
			centerStatusLock.Unlock()
			return serializer.DBErr("无法更新当前记录", err)
		}
		room.Energy += energy
		room.Bill += energy * 5.0
	}
	windSupplyLock.RLock()
	resp := serializer.BuildRoomStatusResponse(room, centerWorkMode, waitStatus[room.RoomID], StartTime)
	windSupplyLock.RUnlock()
	centerStatusLock.Unlock()
	resp.Msg = "房间当前温度更新成功"
	return resp
}
