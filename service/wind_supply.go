package service

import (
	"centralac/model"
	"centralac/serializer"
	"fmt"
	"time"
)

// windSupply 送风函数
func windSupply(room *model.Room) serializer.Response {
	curTime := time.Now()
	record := model.Record{
		RoomID:    room.RoomID,
		StartTime: curTime,
		EndTime:   curTime,
		StartTemp: room.CurrentTemp,
		EndTemp:   room.CurrentTemp,
		WindSpeed: room.WindSpeed,
	}
	if err := model.DB.Create(&record).Error; err != nil {
		return serializer.DBErr("开始送风失败", err)
	}
	model.DB.Where("room_id = ?", room.RoomID).Last(&record)

	roomNew := make(map[string]interface{})
	roomNew["wind_supply"] = true
	roomNew["current_record"] = record.ID
	if err := model.DB.Model(&room).Updates(roomNew).Error; err != nil {
		return serializer.DBErr("房间记录修改失败", err)
	}
	room.WindSupply = true
	room.CurrentRecord = record.ID
	resp := serializer.BuildRoomResponse(*room)
	resp.Msg = "送风成功"
	return resp
}

// stopWindSupply 停止送风
func stopWindSupply(room *model.Room) serializer.Response {
	var record model.Record
	if err := model.DB.Where("id = ?", room.CurrentRecord).First(&record).Error; err != nil {
		return serializer.DBErr("送风记录查找失败", err)
	}
	recordNew := make(map[string]interface{})
	endTime := time.Now()
	recordNew["end_time"] = endTime
	recordNew["end_temp"] = room.CurrentTemp
	minDur := float32(endTime.Sub(record.StartTime).Minutes())
	var energy float32
	switch room.WindSpeed {
	case model.High:
		energy = minDur * 1.2
	case model.Medium:
		energy = minDur
	case model.Low:
		energy = minDur * 0.8
	}
	recordNew["energy"] = energy
	recordNew["bill"] = energy * 5.0
	if err := model.DB.Model(&record).Updates(recordNew).Error; err != nil {
		return serializer.DBErr("送风记录修改失败", err)
	}
	room.WindSupply = false
	room.Energy += energy
	room.Bill += energy * 5.0
	roomNew := make(map[string]interface{})
	roomNew["wind_supply"] = false
	roomNew["energy"] = room.Energy
	roomNew["bill"] = room.Bill
	if err := model.DB.Model(&room).Updates(roomNew).Error; err != nil {
		return serializer.DBErr("房间记录修改失败", err)
	}
	resp := serializer.BuildRoomResponse(*room)
	resp.Msg = "停止送风成功"
	return resp
}

// WindSupplySchedule 送风调度函数
func WindSupplySchedule() {
	ticker := time.NewTicker(10 * time.Second)
	for ; true; <-ticker.C {
		centerStatusLock.Lock()
		windSupplyLock.Lock()
		if waitList.Len() != 0 {
			var room model.Room
			model.DB.Where("room_id = ?", activeList[0]).First(&room)
			activeList = activeList[1:]
			stopWindSupply(&room)
			waitStatus[room.RoomID] = true
			waitList.PushBack(room.RoomID)
			roomID := waitList.Front().Value
			fmt.Print("scheduler schedule")
			fmt.Print(roomID)
			waitList.Remove(waitList.Front())
			delete(waitStatus, roomID.(string))
			model.DB.Where("room_id = ?", roomID).First(&room)
			activeList = append(activeList, room.RoomID)
			fmt.Print(room.RoomID)
			windSupply(&room)
		}
		windSupplyLock.Unlock()
		centerStatusLock.Unlock()
	}
}
