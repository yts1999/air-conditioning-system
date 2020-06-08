package serializer

import "centralac/model"

// Record 温控记录序列化器
type Record struct {
	ID        uint    `json:"id"`
	RoomID    string  `json:"room_id"`
	StartTime int64   `json:"start_time"`
	EndTime   int64   `json:"end_time"`
	StartTemp float32 `json:"start_temp"`
	EndTemp   float32 `json:"end_temp"`
	WindSpeed uint    `json:"wind_speed"`
	Energy    float32 `json:"energy"`
	Bill      float32 `json:"bill"`
}

//Records 多条温控记录序列化器
type Records struct {
	Record []Record `json:"record"`
}

// BuildRecord 序列化温控记录
func BuildRecord(record model.Record) Record {
	return Record{
		ID:        record.ID,
		RoomID:    record.RoomID,
		StartTime: record.StartTime.Unix(),
		EndTime:   record.EndTime.Unix(),
		StartTemp: record.StartTemp,
		EndTemp:   record.EndTemp,
		WindSpeed: record.WindSpeed,
		Energy:    record.Energy,
		Bill:      record.Bill,
	}
}

// BuildRecordResponse 序列化温控记录响应
func BuildRecordResponse(record model.Record) Response {
	return Response{
		Data: BuildRecord(record),
	}
}

// BuildRecordList 序列化温控记录列表
func BuildRecordList(rs []model.Record) []Record {
	var records []Record
	for _, r := range rs {
		record := BuildRecord(r)
		records = append(records, record)
	}
	return records
}

//BuildRecords 序列化多条温控记录
func BuildRecords(records []model.Record) Records {
	return Records{
		Record: BuildRecordList(records),
	}
}

// BuildRecordsResponse 序列化多条温控记录响应
func BuildRecordsResponse(records []model.Record) Response {
	return Response{
		Data: BuildRecords(records),
	}
}
