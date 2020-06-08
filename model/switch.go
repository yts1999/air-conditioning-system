package model

import "time"

// Switch 开关机记录模型
type Switch struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	RoomID string
	Time   time.Time
}

// GetSwitchTimeOfRoom 用房间号获取房间开关机次数
func GetSwitchTimeOfRoom(RoomID interface{}) (int, error) {
	var switches []Switch
	result := DB.Where("room_id = ?", RoomID).Find(&switches)
	return len(switches), result.Error
}

// GetDaySwitchTimeOfRoom 用房间号获取房间指定日开关机次数
func GetDaySwitchTimeOfRoom(RoomID interface{}, Year interface{}, Month interface{}, Day interface{}) (int, error) {
	StartTime := time.Date(Year.(int), time.Month(Month.(int)), Day.(int), 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year.(int), time.Month(Month.(int)), Day.(int), 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}

// GetMonthSwitchTimeOfRoom 用房间号获取房间指定月开关机次数
func GetMonthSwitchTimeOfRoom(RoomID interface{}, Year interface{}, Month interface{}) (int, error) {
	StartTime := time.Date(Year.(int), time.Month(Month.(int)), 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year.(int), time.Month(Month.(int)), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}

// GetYearSwitchTimeOfRoom 用房间号获取房间指定年开关机次数
func GetYearSwitchTimeOfRoom(RoomID interface{}, Year interface{}) (int, error) {
	StartTime := time.Date(Year.(int), 1, 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year.(int), 1, 1, 0, 0, 0, 0, time.Local).AddDate(1, 0, 0).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}
