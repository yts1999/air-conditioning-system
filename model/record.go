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
	WindSpeed uint
	Energy    float32
	Bill      float32
}

// GetRecord 获取房间温控记录
func GetRecord() ([]Record, error) {
	var records []Record
	result := DB.Find(&records)
	return records, result.Error
}

// GetDayRecord 获取指定日温控记录
func GetDayRecord(Year int, Month int, Day int) ([]Record, error) {
	StartTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("start_time <= ? AND end_time >= ?", EndTime, StartTime).Find(&records)
	return records, result.Error
}

// GetWeekRecord 获取指定周温控记录
func GetWeekRecord(Year int, Month int, Day int) ([]Record, error) {
	StartTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, -7)
	EndTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("start_time <= ? AND end_time >= ?", EndTime, StartTime).Find(&records)
	return records, result.Error
}

// GetMonthRecord 获取指定月温控记录
func GetMonthRecord(Year int, Month int) ([]Record, error) {
	StartTime := time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("start_time <= ? AND end_time >= ?", EndTime, StartTime).Find(&records)
	return records, result.Error
}

// GetYearRecord 获取指定年温控记录
func GetYearRecord(Year int) ([]Record, error) {
	StartTime := time.Date(Year, 1, 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, 1, 1, 0, 0, 0, 0, time.Local).AddDate(1, 0, 0).Add(-time.Nanosecond)
	var records []Record
	result := DB.Where("start_time <= ? AND end_time >= ?", EndTime, StartTime).Find(&records)
	return records, result.Error
}

// GetRecordOfRoom 用房间号获取房间温控记录
func GetRecordOfRoom(RoomID string) ([]Record, float32, float32, error) {
	var records []Record
	if err := DB.Where("room_id = ?", RoomID).Find(&records).Error; err != nil {
		return records, 0, 0, err
	}
	// SumResult 求和返回类型
	type SumResult struct {
		Sum float32
	}
	var sumResult SumResult
	if err := DB.Table("records").Select("sum(energy) as sum").Where("room_id = ?", RoomID).Scan(&sumResult).Error; err != nil {
		return records, 0, 0, err
	}
	totalEnergy := sumResult.Sum
	if err := DB.Table("records").Select("sum(bill) as sum").Where("room_id = ?", RoomID).Scan(&sumResult).Error; err != nil {
		return records, totalEnergy, 0, err
	}
	totalBill := sumResult.Sum
	return records, totalEnergy, totalBill, nil
}

// GetDayRecordOfRoom 用房间号获取房间指定日温控记录
func GetDayRecordOfRoom(RoomID string, Year int, Month int, Day int) ([]Record, float32, float32, error) {
	StartTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var records []Record
	if err := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records).Error; err != nil {
		return records, 0, 0, err
	}
	// SumResult 求和返回类型
	type SumResult struct {
		Sum float32
	}
	var sumResult SumResult
	if err := DB.Table("records").Select("sum(energy) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, 0, 0, err
	}
	totalEnergy := sumResult.Sum
	if err := DB.Table("records").Select("sum(bill) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, totalEnergy, 0, err
	}
	totalBill := sumResult.Sum
	return records, totalEnergy, totalBill, nil
}

// GetWeekRecordOfRoom 用房间号获取房间指定周温控记录
func GetWeekRecordOfRoom(RoomID string, Year int, Month int, Day int) ([]Record, float32, float32, error) {
	StartTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, -7)
	EndTime := time.Date(Year, time.Month(Month), Day, 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(-time.Nanosecond)
	var records []Record
	if err := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records).Error; err != nil {
		return records, 0, 0, err
	}
	// SumResult 求和返回类型
	type SumResult struct {
		Sum float32
	}
	var sumResult SumResult
	if err := DB.Table("records").Select("sum(energy) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, 0, 0, err
	}
	totalEnergy := sumResult.Sum
	if err := DB.Table("records").Select("sum(bill) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, totalEnergy, 0, err
	}
	totalBill := sumResult.Sum
	return records, totalEnergy, totalBill, nil
}

// GetMonthRecordOfRoom 用房间号获取房间指定月温控记录
func GetMonthRecordOfRoom(RoomID string, Year int, Month int) ([]Record, float32, float32, error) {
	StartTime := time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, time.Month(Month), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(-time.Nanosecond)
	var records []Record
	if err := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records).Error; err != nil {
		return records, 0, 0, err
	}
	// SumResult 求和返回类型
	type SumResult struct {
		Sum float32
	}
	var sumResult SumResult
	if err := DB.Table("records").Select("sum(energy) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, 0, 0, err
	}
	totalEnergy := sumResult.Sum
	if err := DB.Table("records").Select("sum(bill) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, totalEnergy, 0, err
	}
	totalBill := sumResult.Sum
	return records, totalEnergy, totalBill, nil
}

// GetYearRecordOfRoom 用房间号获取房间指定年温控记录
func GetYearRecordOfRoom(RoomID string, Year int) ([]Record, float32, float32, error) {
	StartTime := time.Date(Year, 1, 1, 0, 0, 0, 0, time.Local)
	EndTime := time.Date(Year, 1, 1, 0, 0, 0, 0, time.Local).AddDate(1, 0, 0).Add(-time.Nanosecond)
	var records []Record
	if err := DB.Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Find(&records).Error; err != nil {
		return records, 0, 0, err
	}
	// SumResult 求和返回类型
	type SumResult struct {
		Sum float32
	}
	var sumResult SumResult
	if err := DB.Table("records").Select("sum(energy) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, 0, 0, err
	}
	totalEnergy := sumResult.Sum
	if err := DB.Table("records").Select("sum(bill) as sum").Where("room_id = ? AND start_time <= ? AND end_time >= ?", RoomID, EndTime, StartTime).Scan(&sumResult).Error; err != nil {
		return records, totalEnergy, 0, err
	}
	totalBill := sumResult.Sum
	return records, totalEnergy, totalBill, nil
}
