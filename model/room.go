package model

// Room 房间模型
type Room struct {
	RoomID     string `gorm:"primary_key"`
	SwitchTime uint
}

// GetRoom 用RoomID获取房间
func GetRoom(RoomID interface{}) (Room, error) {
	var room Room
	result := DB.Where("room_id = ?", RoomID).First(&room)
	return room, result.Error
}