package serializer

import (
	"centralac/model"
	"time"
)

// RoomStatus 房间状态序列化器
type RoomStatus struct {
	RoomID      string  `json:"room_id"`
	PowerOn     bool    `json:"power_on"`
	WindSupply  bool    `json:"wind_supply"`
	CurrentTemp float32 `json:"current_temp"`
	TargetTemp  float32 `json:"target_temp"`
	WindSpeed   uint    `json:"wind_speed"`
	Energy      float32 `json:"energy"`
	Bill        float32 `json:"bill"`
	WorkMode    uint    `json:"work_mode"`
	Wait        bool    `json:"wait"`
	StartTime   int64   `json:"start_time"`
}

// BuildRoomStatus 序列化房间状态
func BuildRoomStatus(room model.Room, workMode uint, waitStatus bool, startTime time.Time) RoomStatus {
	return RoomStatus{
		RoomID:      room.RoomID,
		PowerOn:     room.PowerOn,
		WindSupply:  room.WindSupply,
		CurrentTemp: room.CurrentTemp,
		TargetTemp:  room.TargetTemp,
		WindSpeed:   room.WindSpeed,
		Energy:      room.Energy,
		Bill:        room.Bill,
		WorkMode:    workMode,
		Wait:        waitStatus,
		StartTime:   startTime.Unix(),
	}
}

// BuildRoomStatusResponse 序列化房间状态响应
func BuildRoomStatusResponse(room model.Room, workMode uint, waitStatus bool, startTime time.Time) Response {
	return Response{
		Data: BuildRoomStatus(room, workMode, waitStatus, startTime),
	}
}
