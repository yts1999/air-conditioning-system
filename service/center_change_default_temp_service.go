package service

import (
	"centralac/serializer"
)

// CenterChangeDefaultTempService 中央空调修改默认温度的服务
type CenterChangeDefaultTempService struct {
	Temp float32 `form:"temp" json:"temp" binding:"required"`
}

// Change 中央空调修改默认温度函数
func (service *CenterChangeDefaultTempService) Change() serializer.Response {
	centerStatusLock.Lock()
	if service.Temp > highestTemp || service.Temp < lowestTemp {
		centerStatusLock.Unlock()
		return serializer.SystemErr("温度超出范围", nil)
	}
	defaultTemp = service.Temp
	windSupplyLock.RLock()
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	windSupplyLock.RUnlock()
	centerStatusLock.Unlock()
	resp.Msg = "中央空调修改默认温度成功"
	return resp
}
