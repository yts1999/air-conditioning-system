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
var centerWorkMode uint = 1 // 1 -- 制冷 2 -- 制热
var activeList = []string{}
var defaultTemp float32 = 22.0
var lowestTemp float32 = 18.0
var highestTemp float32 = 25.0

// Startup 中央空调开机函数
func (service *CenterStartupService) Startup() serializer.Response {
	centerStatusLock.Lock()
	windSupplyLock.Lock()
	centerPowerOn = true
	centerWorkMode = 1
	activeList = activeList[0:0]
	defaultTemp = 22.0
	lowestTemp = 18.0
	highestTemp = 25.0
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	windSupplyLock.Unlock()
	centerStatusLock.Unlock()
	resp.Msg = "中央空调开机成功"
	return resp
}
