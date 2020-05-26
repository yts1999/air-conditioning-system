package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RecordOfRoomOfMonthService 获取房间月记录的服务
type RecordOfRoomOfMonthService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Year   uint   `form:"year" json:"year" binding:"required"`
	Month  uint   `form:"month" json:"month" binding:"required,gte=1,lte=12"`
}

// List 获取房间月记录函数
func (service *RecordOfRoomOfMonthService) List() serializer.Response {
	records, err := model.GetMonthRecordOfRoom(service.RoomID, service.Year, service.Month)
	if err != nil {
		return serializer.DBErr("获取房间月记录失败", err)
	}
	resp := serializer.BuildRecordsResponse(records)
	resp.Msg = "获取房间月记录成功"
	return resp
}
