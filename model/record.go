package model

import "time"

// Record 温控记录模型
type Record struct {
	ID        uint `gorm:"primary_key;auto_increment"`
	RoomID    string
	StartTime time.Time
	EndTime   time.Time
	StartTemp float32
	EndTemp   float32
	Energy    float32
	Bill      float32
}

// GetRecordOfRoom 用房间号获取房间温控记录
func GetRecordOfRoom(RoomID interface{}) ([]Record, error) {
	var records []Record
	result := DB.Where("room_id = ?", RoomID).Find(&records)
	return records, result.Error
}

// GetDayRecordOfRoom 用房间号获取房间指定日温控记录
func GetDayRecordOfRoom(RoomID interface{}, Year interface{}, Month interface{}, Day interface{}) ([]Record, error) {
	StartTime := time.Date(Year.(int), time.Month(Month.(int)), Day.(int), 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year.(int), time.Month(Month.(int)), Day.(int), 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records)
	return records, result.Error
}

// GetMonthRecordOfRoom 用房间号获取房间指定月温控记录
func GetMonthRecordOfRoom(RoomID interface{}, Year interface{}, Month interface{}) ([]Record, error) {
	StartTime := time.Date(Year.(int), time.Month(Month.(int)), 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year.(int), time.Month(Month.(int)), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records)
	return records, result.Error
}

// GetYearRecordOfRoom 用房间号获取房间指定年温控记录
func GetYearRecordOfRoom(RoomID interface{}, Year interface{}) ([]Record, error) {
	StartTime := time.Date(Year.(int), 1, 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year.(int), 1, 1, 0, 0, 0, 0, time.Local).AddDate(1, 0, 0).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records)
	return records, result.Error
}
