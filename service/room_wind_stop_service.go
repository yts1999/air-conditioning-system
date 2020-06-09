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

	centerStatusLock.Lock()
	if !centerPowerOn {
		centerStatusLock.Unlock()
		return serializer.SystemErr("中央空调未开启", nil)
	}

	for i := 0; i < len(activeList); i++ {
		if activeList[i] == room.RoomID {
			activeList = append(activeList[:i], activeList[i+1:]...)
			break
		}
	}
	resp := stopWindSupply(&room)
	if resp.Code != 0 {
		centerStatusLock.Unlock()
		return resp
	}

	windSupplyLock.Lock()
	if waitList.Len() != 0 {
		roomID := waitList.Front().Value
		waitList.Remove(waitList.Front())
		delete(waitStatus, roomID.(string))
		var windRoom model.Room
		model.DB.Where("room_id = ?", roomID).First(&windRoom)
		windSupplyLock.Unlock()
		activeList = append(activeList, windRoom.RoomID)
		resp := windSupply(&windRoom)
		if resp.Code != 0 {
			centerStatusLock.Unlock()
			return resp
		}
	} else {
		windSupplySem++
		windSupplyLock.Unlock()
	}
	centerStatusLock.Unlock()
	return resp
}
