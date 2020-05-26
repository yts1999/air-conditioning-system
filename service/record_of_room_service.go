package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RecordOfRoomService 获取房间所有记录的服务
type RecordOfRoomService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// List 获取房间所有记录函数
func (service *RecordOfRoomService) List() serializer.Response {
	records, err := model.GetRecordOfRoom(service.RoomID)
	if err != nil {
		return serializer.DBErr("获取房间所有记录失败", err)
	}
	resp := serializer.BuildRecordsResponse(records)
	resp.Msg = "获取房间所有记录成功"
	return resp
}
