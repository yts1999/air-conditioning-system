package service

import (
	"centralac/model"
	"centralac/serializer"
	"fmt"
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
		windSupplyLock.Lock()
		for i := 0; i < len(activeList); i++ {
			if activeList[i] == room.RoomID {
				activeList = append(activeList[:i], activeList[i+1:]...)
				break
			}
		}
		resp := stopWindSupply(&room)
		if resp.Code != 0 {
			centerStatusLock.Unlock()
			return resp
		}
		if centerPowerOn {
			if waitList.Len() != 0 {
				roomID := waitList.Front().Value
				waitList.Remove(waitList.Front())
				delete(waitStatus, roomID.(string))
				var windRoom model.Room
				model.DB.Where("room_id = ?", roomID).First(&windRoom)
				activeList = append(activeList, windRoom.RoomID)
				resp := windSupply(&windRoom)
				if resp.Code != 0 {
					windSupplyLock.Unlock()
					centerStatusLock.Unlock()
					return resp
				}
			}
		}
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
	} else {
		windSupplyLock.Lock()
		fmt.Print("shutdown waitlist")
		for i := waitList.Front(); i != nil; i = i.Next() {
			if i.Value == room.RoomID {
				fmt.Print(room.RoomID)
				waitList.Remove(i)
				delete(waitStatus, room.RoomID)
				fmt.Print(waitStatus[room.RoomID])
				break
			}
		}
		windSupplyLock.Unlock()
	}

	// 从控机关机
	if err := model.DB.Model(&room).Update("power_on", false).Error; err != nil {
		return serializer.DBErr("从控机关机失败", err)
	}
	room.PowerOn = false

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "从控机关机成功"
	return resp
}
