package serializer

import "centralac/model"

// Room 房间序列化器
type Room struct {
	RoomID      string  `json:"room_id"`
	SwitchTime  uint    `json:"switch_time"`
	PowerOn     bool    `json:"power_on"`
	CurrentTemp float32 `json:"current_temp"`
	TargetTemp  float32 `json:"target_temp"`
	WindSpeed   uint    `json:"wind_speed"`
}

// BuildRoom 序列化房间
func BuildRoom(room model.Room) Room {
	return Room{
		RoomID:      room.RoomID,
		SwitchTime:  room.SwitchTime,
		PowerOn:     room.PowerOn,
		CurrentTemp: room.CurrentTemp,
		TargetTemp:  room.TargetTemp,
		WindSpeed:   room.WindSpeed,
	}
}

// BuildRoomResponse 序列化房间响应
func BuildRoomResponse(room model.Room) Response {
	return Response{
		Data: BuildRoom(room),
	}
}
