package serializer

import "centralac/model"

// Room 房间序列化器
type Room struct {
	RoomID      string  `json:"room_id"`
	SwitchTime  uint    `json:"switch_time"`
	PowerOn     bool    `json:"power_on"`
	WindSupply  bool    `json:"wind_supply"`
	CurrentTemp float32 `json:"current_temp"`
	TargetTemp  float32 `json:"target_temp"`
	WindSpeed   uint    `json:"wind_speed"`
	Energy      float32 `json:"energy"`
	Bill        float32 `json:"bill"`
}

// BuildRoom 序列化房间
func BuildRoom(room model.Room) Room {
	return Room{
		RoomID:      room.RoomID,
		SwitchTime:  room.SwitchTime,
		PowerOn:     room.PowerOn,
		WindSupply:  room.WindSupply,
		CurrentTemp: room.CurrentTemp,
		TargetTemp:  room.TargetTemp,
		WindSpeed:   room.WindSpeed,
		Energy:      room.Energy,
		Bill:        room.Bill,
	}
}

// BuildRoomResponse 序列化房间响应
func BuildRoomResponse(room model.Room) Response {
	return Response{
		Data: BuildRoom(room),
	}
}

// BuildRooms 序列化所有房间
func BuildRooms(rs []model.Room) (rooms []Room) {
	for _, r := range rs {
		room := BuildRoom(r)
		rooms = append(rooms, room)
	}
	return rooms
}

// BuildRoomsResponse 序列化所有房间响应
func BuildRoomsResponse(rooms []model.Room) Response {
	return Response{
		Data: BuildRooms(rooms),
	}
}
