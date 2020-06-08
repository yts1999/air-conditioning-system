package model

// Room 房间模型
type Room struct {
	RoomID        string `gorm:"primary_key"`
	SwitchTime    uint
	PowerOn       bool
	WindSupply    bool
	CurrentTemp   float32
	TargetTemp    float32
	WindSpeed     uint
	CurrentRecord uint `sql:"default:null"`
	Energy        float32
	Bill          float32
}

// 风速
const (
	Low uint = iota + 1
	Medium
	High
)

// GetRoom 用RoomID获取房间
func GetRoom(RoomID interface{}) (Room, error) {
	var room Room
	result := DB.Where("room_id = ?", RoomID).First(&room)
	return room, result.Error
}
