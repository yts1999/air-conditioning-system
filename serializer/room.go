package serializer

import "centralac/model"

// Room 房间序列化器
type Room struct {
	RoomID     string `json:"room_id"`
	SwitchTime uint   `json:"switch_time"`
}

// BuildRoom 序列化房间
func BuildRoom(room model.Room) Room {
	return Room{
		RoomID:     room.RoomID,
		SwitchTime: room.SwitchTime,
	}
}

// BuildRoomResponse 序列化房间响应
func BuildRoomResponse(room model.Room) Response {
	return Response{
		Data: BuildRoom(room),
	}
}
