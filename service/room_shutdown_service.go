package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RoomShutdownService 从控机关机的服务
type RoomShutdownService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Shutdown 从控机关机函数
func (service *RoomShutdownService) Shutdown() serializer.Response {
	var room model.Room
	if err := model.DB.First(&room, service.RoomID).Error; err != nil {
		return serializer.Err(404, "房间信息不存在", err)
	}

	//停止送风
	if room.WindSupply {
		centerStatusLock.Lock()
		resp := stopWindSupply(&room)
		if resp.Code != 0 {
			centerStatusLock.Unlock()
			return resp
		}
		if centerPowerOn {
			windSupplyLock.Lock()
			waitListLock.Lock()
			if waitList.Len() != 0 {
				roomID := waitList.Front().Value
				waitList.Remove(waitList.Front())
				delete(waitStatus, roomID.(string))
				waitListLock.Unlock()
				var windRoom model.Room
				model.DB.Where("room_id = ?", roomID).First(&windRoom)
				windSupplyLock.Unlock()
				resp := windSupply(&windRoom)
				if resp.Code != 0 {
					centerStatusLock.Unlock()
					return resp
				}
			} else {
				waitListLock.Unlock()
				windSupplySem++
				windSupplyLock.Unlock()
			}
		}
		centerStatusLock.Unlock()
	} else {
		waitListLock.Lock()
		for i := waitList.Front(); i != nil; i = i.Next() {
			if i.Value == room.RoomID {
				waitList.Remove(i)
				delete(waitStatus, room.RoomID)
				break
			}
		}
		waitListLock.Unlock()
	}

	// 从控机关机
	err := model.DB.Model(model.Room{}).
		Where("room_id = ?", service.RoomID).
		Update("power_on", false).Error
	if err != nil {
		return serializer.DBErr("从控机关机失败", err)
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "从控机关机成功"
	return resp
}
