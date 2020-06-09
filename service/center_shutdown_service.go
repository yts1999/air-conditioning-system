package service

import (
	"centralac/model"
	"centralac/serializer"
	"fmt"
)

// CenterShutdownService 中央空调关机的服务
type CenterShutdownService struct {
}

// Shutdown 中央空调关机函数
func (service *CenterShutdownService) Shutdown() serializer.Response {
	centerStatusLock.Lock()
	fmt.Printf("starting shutdown\n")
	centerPowerOn = false
	fmt.Printf("%d", len(activeList))
	roomList := activeList
	activeList = activeList[0:0]
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	windSupplyLock.Lock()
	windSupplySem = 3
	waitListLock.Lock()
	fmt.Printf("wait list\n")
	waitList.Init()
	waitStatus = make(map[string]bool)
	waitListLock.Unlock()
	for i := 0; i < len(roomList); i++ {
		if roomList[i] != "" {
			fmt.Printf("%s", roomList[i])
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
