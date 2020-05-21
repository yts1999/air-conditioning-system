package serializer

import "centralac/model"

// Guest 房客序列化器
type Guest struct {
	ID        uint   `json:"id"`
	RoomID    string `json:"room_id"`
	CreatedAt int64  `json:"created_at"`
}

// BuildGuest 序列化房客
func BuildGuest(guest model.Guest) Guest {
	return Guest{
		ID:        guest.ID,
		RoomID:    guest.RoomID,
		CreatedAt: guest.CreatedAt.Unix(),
	}
}

// BuildGuestResponse 序列化房客响应
func BuildGuestResponse(guest model.Guest) Response {
	return Response{
		Data: BuildGuest(guest),
	}
}
