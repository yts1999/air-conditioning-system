package service

import (
	"centralac/model"
	"centralac/serializer"
)

// GuestDeleteService 管理删除房客的服务
type GuestDeleteService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Delete 删除房客函数
func (service *GuestDeleteService) Delete() serializer.Response {
	//检查房间是否存在
	var room model.Room
	if !model.DB.Where("room_id = ?", service.RoomID).First(&room).RecordNotFound() {
		return serializer.ParamErr("房间不存在", nil)
	}

	//停止送风
	if room.WindSupply {
		centerStatusLock.Lock()
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
			windSupplyLock.Lock()
			if waitList.Len() != 0 {
				roomID := waitList.Front().Value
				waitList.Remove(waitList.Front())
				delete(waitStatus, roomID.(string))
				var windRoom model.Room
				model.DB.Where("room_id = ?", roomID).First(&windRoom)
				windSupplyLock.Unlock()
				activeList = append(activeList, windRoom.RoomID)
				resp := windSupply(&windRoom)
				if resp.Code != 0 {
					centerStatusLock.Unlock()
					return resp
				}
			} else {
				windSupplySem++
				windSupplyLock.Unlock()
			}
		}
		centerStatusLock.Unlock()
	} else {
		windSupplyLock.Lock()
		for i := waitList.Front(); i != nil; i = i.Next() {
			if i.Value == room.RoomID {
				waitList.Remove(i)
				delete(waitStatus, room.RoomID)
				break
			}
		}
		windSupplyLock.Unlock()
	}

	//检查房客是否存在
	var guest model.Guest
	if !model.DB.Where("room_id = ?", service.RoomID).First(&guest).RecordNotFound() {
		return serializer.ParamErr("房客不存在", nil)
	}

	// 删除房客
	if err := model.DB.Where("room_id = ?", service.RoomID).Delete(&guest).Error; err != nil {
		return serializer.ParamErr("房客删除失败", err)
	}

	//清空房间信息
	roomNew := make(map[string]interface{})
	roomNew["power_on"] = false
	centerStatusLock.RLock()
	roomNew["target_temp"] = defaultTemp
	centerStatusLock.RUnlock()
	roomNew["wind_speed"] = model.Medium
	roomNew["energy"] = 0
	roomNew["bill"] = 0
	if err := model.DB.Model(&room).Updates(roomNew).Error; err != nil {
		return serializer.DBErr("房间消费清空失败", err)
	}

	resp := serializer.BuildRoomResponse(room)
	resp.Msg = "房客删除成功"
	return resp
}
