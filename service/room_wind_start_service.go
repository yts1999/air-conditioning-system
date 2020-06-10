package service

import (
	"centralac/model"
	"centralac/serializer"
	"container/list"
	"fmt"
	"sync"
)

// RoomWindStartService 请求送风的服务
type RoomWindStartService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

var windSupplyLock sync.RWMutex
var waitList = list.New()
var waitStatus map[string]bool

// Start 请求送风函数
func (service *RoomWindStartService) Start() serializer.Response {
	var room model.Room
	if model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间号不存在", nil)
	}

	if room.WindSupply {
		return serializer.SystemErr("当前已在送风", nil)
	}

	centerStatusLock.Lock()
	if !centerPowerOn {
		centerStatusLock.Unlock()
		return serializer.SystemErr("中央空调未开启", nil)
	}

	if (room.TargetTemp > room.CurrentTemp && centerWorkMode == 1) || (room.TargetTemp < room.CurrentTemp && centerWorkMode == 2) {
		centerStatusLock.Unlock()
		return serializer.SystemErr("冷暖模式与中央空调不符", nil)
	}

	windSupplyLock.Lock()
	if len(activeList) < 3 {
		//开始送风
		activeList = append(activeList, room.RoomID)
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
		return windSupply(&room)
	}
	fmt.Printf("123\n")
	if !waitStatus[room.RoomID] {
		fmt.Printf("111\n")
		waitList.PushBack(room.RoomID)
		waitStatus[room.RoomID] = true
	}
	windSupplyLock.Unlock()
	centerStatusLock.Unlock()
	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "送风阻塞"
	return resp
}
