package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomWindStopService 停止送风的服务
type RoomWindStopService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Stop 停止送风函数
func (service *RoomWindStopService) Stop() serializer.Response {
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}

	if !room.WindSupply {
		return serializer.SystemErr("当前未在送风", nil)
	}

	centerStatusLock.RLock()
	if !centerPowerOn {
		centerStatusLock.RUnlock()
		return serializer.SystemErr("中央空调未开启", nil)
	}

	resp := stopWindSupply(&room)
	if resp.Code != 0 {
		centerStatusLock.RUnlock()
		return resp
	}

	windSupplyLock.Lock()
	waitListLock.Lock()
	if waitList.Len() != 0 {
		roomID := waitList.Front().Value
		waitList.Remove(waitList.Front())
		delete(waitStatus, roomID.(string))
		waitListLock.Unlock()
		var windRoom model.Room
		model.DB.Where("room_id = ?", roomID).First(&windRoom)
		windSupplyLock.Unlock()
		resp := windSupply(&windRoom)
		if resp.Code != 0 {
			centerStatusLock.RUnlock()
			return resp
		}
	} else {
		waitListLock.Unlock()
		windSupplySem++
		windSupplyLock.Unlock()
	}
	centerStatusLock.RUnlock()
	return resp
}
