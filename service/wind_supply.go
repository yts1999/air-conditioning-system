package service

import (
	"centralac/model"
	"centralac/serializer"
	"time"
)

// windSupply 送风函数
func windSupply(room *model.Room) serializer.Response {
	centerStatusLock.Lock()
	activeList = append(activeList, room.RoomID)
	centerStatusLock.Unlock()
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
	centerStatusLock.Lock()
	for i := 0; i < len(activeList); i++ {
		if activeList[i] == room.RoomID {
			activeList = append(activeList[:i], activeList[:i+1]...)
			break
		}
	}
	centerStatusLock.Unlock()
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
