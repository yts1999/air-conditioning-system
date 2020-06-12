package service

import (
	"centralac/model"
	"centralac/serializer"
	"container/list"
	"sync"
)

// RoomWindStartService 请求送风的服务
type RoomWindStartService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

var windSupplyLock sync.RWMutex
var waitList = list.New()
var waitStatus = make(map[string]bool)

// Start 请求送风函数
func (service *RoomWindStartService) Start() serializer.Response {
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}

	centerStatusLock.Lock()
	windSupplyLock.Lock()

	if !room.PowerOn {
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return serializer.SystemErr("从控机未开启", nil)
	}

	if room.WindSupply {
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return serializer.SystemErr("当前已在送风", nil)
	}

	if !centerPowerOn {
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return serializer.SystemErr("中央空调未开启", nil)
	}

	if (room.TargetTemp > room.CurrentTemp && centerWorkMode == 1) || (room.TargetTemp < room.CurrentTemp && centerWorkMode == 2) {
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return serializer.SystemErr("冷暖模式与中央空调不符", nil)
	}

	if len(activeList) < 3 {
		//开始送风
		activeList = append(activeList, room.RoomID)
		resp := windSupply(&room)
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return resp
	}
	if !waitStatus[room.RoomID] {
		waitList.PushBack(room.RoomID)
		waitStatus[room.RoomID] = true
	}
	windSupplyLock.Unlock()
	centerStatusLock.Unlock()
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "送风阻塞"
	return resp
}
