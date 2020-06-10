package service

import (
	"centralac/serializer"
)

// CenterStatusService 获取中央空调状态的服务
type CenterStatusService struct {
}

// Status 中央空调状态函数
func (service *CenterStatusService) Status() serializer.Response {
	centerStatusLock.RLock()
	windSupplyLock.RLock()
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	windSupplyLock.RUnlock()
	centerStatusLock.RUnlock()
	resp.Msg = "获取中央空调状态成功"
	return resp
}
