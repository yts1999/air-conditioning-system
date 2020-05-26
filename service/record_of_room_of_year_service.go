package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RecordOfRoomOfYearService 获取房间年记录的服务
type RecordOfRoomOfYearService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	Year   uint   `form:"year" json:"year" binding:"required"`
}

// List 获取房间年记录函数
func (service *RecordOfRoomOfYearService) List() serializer.Response {
	records, err := model.GetYearRecordOfRoom(service.RoomID, service.Year)
	if err != nil {
		return serializer.DBErr("获取房间年记录失败", err)
	}
	resp := serializer.BuildRecordsResponse(records)
	resp.Msg = "获取房间年记录成功"
	return resp
}
