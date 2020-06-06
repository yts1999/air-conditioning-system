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
	centerPowerOn = false
	resp := serializer.BuildCenterResponse(centerPowerOn, centerPowerMode, defaultTemp, lowestTemp, highestTemp)
	windSupplyLock.Lock()
	windSupplySem = 3
	waitListLock.Lock()
	waitList.Init()
	waitListLock.Unlock()
	runningListLock.Lock()
	roomList := runningList
	runningList = [3]string{"", "", ""}
	runningListLock.Unlock()
	for i := 0; i < 3; i++ {
		if roomList[i] != "" {
			var room model.Room
			model.DB.Where("room_id = ?", roomList[i]).First(&room)
			resp := stopWindSupply(&room)
			if resp.Code != 0 {
				windSupplyLock.Unlock()
				centerStatusLock.Unlock()
				return resp
			}
		}
	}
	windSupplyLock.Unlock()
	centerStatusLock.Unlock()
	resp.Msg = "中央空调关机成功"
	return resp
}
