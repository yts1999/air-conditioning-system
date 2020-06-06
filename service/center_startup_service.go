package service

import (
	"centralac/serializer"
	"sync"
)

// CenterStartupService 中央空调开机的服务
type CenterStartupService struct {
}

var centerStatusLock sync.RWMutex
var centerPowerOn bool = false
var centerPowerMode uint = 0 // 0 -- 制冷 1 -- 制热
var defaultTemp float32 = 22.0
var lowestTemp float32 = 18.0
var highestTemp float32 = 25.0

// Startup 中央空调开机函数
func (service *CenterStartupService) Startup() serializer.Response {
	centerStatusLock.Lock()
	centerPowerOn = true
	centerPowerMode = 0
	defaultTemp = 22.0
	lowestTemp = 18.0
	highestTemp = 25.0
	resp := serializer.BuildCenterResponse(centerPowerOn, centerPowerMode, defaultTemp, lowestTemp, highestTemp)
	centerStatusLock.Unlock()
	resp.Msg = "中央空调开机成功"
	return resp
}
