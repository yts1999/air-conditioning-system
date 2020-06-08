package serializer

import "centralac/model"

// RoomRecord 房间记录序列化器
type RoomRecord struct {
	RoomID      string   `json:"room_id"`
	Count       int      `json:"count"`
	TotalEnergy float32  `json:"total_energy"`
	TotalBill   float32  `json:"total_bill"`
	Record      []Record `json:"record"`
}

// BuildRoomRecord 序列化开关机次数
func BuildRoomRecord(room model.Room, count int, totalEnergy float32, totalBill float32, records []model.Record) RoomRecord {
	return RoomRecord{
		RoomID:      room.RoomID,
		Count:       count,
		TotalEnergy: totalEnergy,
		TotalBill:   totalBill,
		Record:      BuildRecordList(records),
	}
}

// BuildRoomRecordResponse 序列化开关机次数响应
func BuildRoomRecordResponse(room model.Room, count int, totalEnergy float32, totalBill float32, records []model.Record) Response {
	return Response{
		Data: BuildRoomRecord(room, count, totalEnergy, totalBill, records),
	}
}
