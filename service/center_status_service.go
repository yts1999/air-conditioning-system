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
	resp := serializer.BuildCenterResponse(centerPowerOn, centerPowerMode, defaultTemp, lowestTemp, highestTemp)
	centerStatusLock.RUnlock()
	resp.Msg = "获取中央空调状态成功"
	return resp
}
