package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RecordOfRoomService 获取房间记录的服务
type RecordOfRoomService struct {
	Type   string `form:"type"`
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Year   uint   `form:"year" json:"year"`
	Month  uint   `form:"month" json:"month" binding:"gte=1,lte=12"`
	Day    uint   `form:"day" json:"day" binding:"gte=1,lte=31"`
}

// List 获取房间记录函数
func (service *RecordOfRoomService) List() serializer.Response {
	var records []model.Record
	var err error
	switch service.Type {
	case "all":
		records, err = model.GetRecordOfRoom(service.RoomID)
	case "day":
		records, err = model.GetDayRecordOfRoom(service.RoomID, service.Year, service.Month, service.Day)
	case "month":
		records, err = model.GetMonthRecordOfRoom(service.RoomID, service.Year, service.Month)
	case "year":
		records, err = model.GetYearRecordOfRoom(service.RoomID, service.Year)
	}
	if err != nil {
		return serializer.DBErr("获取房间记录失败", err)
	}
	resp := serializer.BuildRecordsResponse(records)
	resp.Msg = "获取房间记录成功"
	return resp
}
