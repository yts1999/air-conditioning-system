package service

import (
	"centralac/model"
	"centralac/serializer"
)

// CenterShutdownService 中央空调关机的服务
type CenterShutdownService struct {
}

// Shutdown 中央空调关机函数
func (service *CenterShutdownService) Shutdown() serializer.Response {
	centerStatusLock.Lock()
	windSupplyLock.Lock()
	centerPowerOn = false
	roomList := activeList
	activeList = activeList[0:0]
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	waitList.Init()
	waitStatus = make(map[string]bool)
	for i := 0; i < len(roomList); i++ {
		var room model.Room
		model.DB.Where("room_id = ?", roomList[i]).First(&room)
		stopWindSupply(&room)
	}
	windSupplyLock.Unlock()
	centerStatusLock.Unlock()
	resp.Msg = "中央空调关机成功"
	return resp
}
