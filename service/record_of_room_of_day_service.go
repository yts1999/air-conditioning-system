package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RecordOfRoomOfDayService 获取房间日记录的服务
type RecordOfRoomOfDayService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Year   uint   `form:"year" json:"year" binding:"required"`
	Month  uint   `form:"month" json:"month" binding:"required,gte=1,lte=12"`
	Day    uint   `form:"day" json:"day" binding:"required,gte=1,lte=31"`
}

// List 获取房间日记录函数
func (service *RecordOfRoomOfDayService) List() serializer.Response {
	records, err := model.GetDayRecordOfRoom(service.RoomID, service.Year, service.Month, service.Day)
	if err != nil {
		return serializer.DBErr("获取房间日记录失败", err)
	}
	resp := serializer.BuildRecordsResponse(records)
	resp.Msg = "获取房间日记录成功"
	return resp
}
