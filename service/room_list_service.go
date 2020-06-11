package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomListService 获取所有房间信息的服务
type RoomListService struct {
}

// List 获取所有房间信息函数
func (service *RoomListService) List() serializer.Response {
	rooms := []model.Room{}
	if err := model.DB.Find(&rooms).Error; err != nil {
		return serializer.DBErr("获取所有房间信息失败", err)
	}
	centerStatusLock.RLock()
	windSupplyLock.RLock()
	for i := 0; i < len(rooms); i++ {
		if rooms[i].WindSupply {
			var record model.Record
			if err := model.DB.First(&record, rooms[i].CurrentRecord).Error; err != nil {
				windSupplyLock.RUnlock()
				centerStatusLock.RUnlock()
				return serializer.SystemErr("无法查询当前记录", err)
			}
			rooms[i].Energy += record.Energy
			rooms[i].Bill += record.Bill
		}
	}
	windSupplyLock.RUnlock()
	centerStatusLock.RUnlock()
	resp := serializer.BuildRoomsResponse(rooms)
	resp.Msg = "获取所有房间信息成功"
	return resp
}
