package service

import (
	"centralac/model"
	"centralac/serializer"
	"time"
)

// RoomStartupService 从控机开机的服务
type RoomStartupService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Startup 从控机开机函数
func (service *RoomStartupService) Startup() serializer.Response {
	var room model.Room
	if err := model.DB.First(&room, service.RoomID).Error; err != nil {
		return serializer.Err(404, "房间信息不存在", err)
	}

	roomNew := make(map[string]interface{})
	roomNew["power_on"] = true
	roomNew["target_temp"] = defaultTemp
	roomNew["wind_speed"] = model.Medium
	if err := model.DB.Model(&room).Updates(roomNew).Error; err != nil {
		return serializer.DBErr("从控机开机失败", err)
	}

	switchRecord := model.Switch{
		RoomID: room.RoomID,
		Time:   time.Now(),
	}
	if err := model.DB.Create(&switchRecord).Error; err != nil {
		return serializer.DBErr("从控机开机失败", err)
	}

	centerStatusLock.RLock()
	resp := serializer.BuildCenterResponse(centerPowerOn, centerWorkMode, activeList, defaultTemp, lowestTemp, highestTemp)
	centerStatusLock.RUnlock()
	resp.Msg = "从控机开机成功"
	return resp
}
