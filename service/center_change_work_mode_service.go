package service

import (
	"centralac/serializer"
)

// CenterChangeWorkModeService 中央空调改变工作模式的服务
type CenterChangeWorkModeService struct {
	WorkMode uint `form:"work_mode" json:"work_mode" binding:"required,gte=1,lte=2"`
}

// Change 中央空调改变工作模式函数
func (service *CenterChangeWorkModeService) Change() serializer.Response {
	centerStatusLock.Lock()
	centerWorkMode = service.WorkMode
	if centerWorkMode == 1 {
		defaultTemp = 22.0
		lowestTemp = 18.0
		highestTemp = 25.0
	} else {
		defaultTemp = 28.0
		lowestTemp = 25.0
		highestTemp = 30.0
	}
	windSupplyLock.RLock()
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	windSupplyLock.RUnlock()
	centerStatusLock.Unlock()
	resp.Msg = "中央空调改变工作模式成功"
	return resp
}
