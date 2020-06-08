package model

import "time"

// Switch 开关机记录模型
type Switch struct {
	ID     uint `gorm:"primary_key;auto_increment"`
	RoomID string
	Time   time.Time
}

// GetSwitchTimeOfRoom 用房间号获取房间开关机次数
func GetSwitchTimeOfRoom(RoomID string) (int, error) {
	var switches []Switch
	result := DB.Where("room_id = ?", RoomID).Find(&switches)
	return len(switches), result.Error
}

// GetDaySwitchTimeOfRoom 用房间号获取房间指定日开关机次数
func GetDaySwitchTimeOfRoom(RoomID string, Year int, Month int, Day int) (int, error) {
	StartTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND time <= ? AND time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}

// GetWeekSwitchTimeOfRoom 用房间号获取房间指定周开关机次数
func GetWeekSwitchTimeOfRoom(RoomID string, Year int, Month int, Day int) (int, error) {
	StartTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, -7)
	EndTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND time <= ? AND time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}

// GetMonthSwitchTimeOfRoom 用房间号获取房间指定月开关机次数
func GetMonthSwitchTimeOfRoom(RoomID string, Year int, Month int) (int, error) {
	StartTime := time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND time <= ? AND time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}

// GetYearSwitchTimeOfRoom 用房间号获取房间指定年开关机次数
func GetYearSwitchTimeOfRoom(RoomID string, Year int) (int, error) {
	StartTime := time.Date(Year, 1, 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, 1, 1, 0, 0, 0, 0, time.Local).AddDate(1, 0, 0).Add(-time.Nanosecond)
	var switches []Switch
	result := DB.Where("room_id = ? AND time <= ? AND time >= ?", RoomID, EndTime, StartTime).Find(&switches)
	return len(switches), result.Error
}
