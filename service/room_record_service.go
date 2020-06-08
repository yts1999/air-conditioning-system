package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomRecordService 获取房间记录的服务
type RoomRecordService struct {
	Type   string `form:"type" json:"type" binding:"required"`
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Year   int    `form:"year" json:"year"`
	Month  int    `form:"month" json:"month"`
	Day    int    `form:"day" json:"day"`
}

// List 获取房间记录函数
func (service *RoomRecordService) List() serializer.Response {
	var room model.Room
	if err := model.DB.First(&room, service.RoomID).Error; err != nil {
		return serializer.Err(404, "房间信息不存在", err)
	}

	var records []model.Record
	var totalEnergy float32
	var totalBill float32
	var err error
	switch service.Type {
	case "all":
		records, totalEnergy, totalBill, err = model.GetRecordOfRoom(service.RoomID)
	case "day":
		records, totalEnergy, totalBill, err = model.GetDayRecordOfRoom(service.RoomID, service.Year, service.Month, service.Day)
	case "week":
		records, totalEnergy, totalBill, err = model.GetWeekRecordOfRoom(service.RoomID, service.Year, service.Month, service.Day)
	case "month":
		records, totalEnergy, totalBill, err = model.GetMonthRecordOfRoom(service.RoomID, service.Year, service.Month)
	case "year":
		records, totalEnergy, totalBill, err = model.GetYearRecordOfRoom(service.RoomID, service.Year)
	}
	if err != nil {
		return serializer.DBErr("获取房间记录失败", err)
	}

	var switchTime int
	switch service.Type {
	case "all":
		switchTime, err = model.GetSwitchTimeOfRoom(service.RoomID)
	case "day":
		switchTime, err = model.GetDaySwitchTimeOfRoom(service.RoomID, service.Year, service.Month, service.Day)
	case "week":
		switchTime, err = model.GetWeekSwitchTimeOfRoom(service.RoomID, service.Year, service.Month, service.Day)
	case "month":
		switchTime, err = model.GetMonthSwitchTimeOfRoom(service.RoomID, service.Year, service.Month)
	case "year":
		switchTime, err = model.GetYearSwitchTimeOfRoom(service.RoomID, service.Year)
	}
	if err != nil {
		return serializer.DBErr("获取房间开关机记录失败", err)
	}

	resp := serializer.BuildRoomRecordResponse(room, switchTime, totalEnergy, totalBill, records)
	resp.Msg = "获取房间记录成功"
	return resp
}
