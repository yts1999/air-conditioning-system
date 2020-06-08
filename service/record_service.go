package service

import (
	"centralac/model"
	"centralac/serializer"
)

// RecordService 获取记录的服务
type RecordService struct {
	Type  string `form:"type" json:"type" binding:"required"`
	Year  uint   `form:"year" json:"year"`
	Month uint   `form:"month" json:"month"`
	Day   uint   `form:"day" json:"day"`
}

// List 获取房间记录函数
func (service *RecordService) List() serializer.Response {
	var records []model.Record
	var err error
	switch service.Type {
	case "all":
		records, err = model.GetRecord()
	case "day":
		records, err = model.GetDayRecord(service.Year, service.Month, service.Day)
	case "month":
		records, err = model.GetMonthRecord(service.Year, service.Month)
	case "year":
		records, err = model.GetYearRecord(service.Year)
	}
	if err != nil {
		return serializer.DBErr("获取记录失败", err)
	}
	resp := serializer.BuildRecordsResponse(records)
	resp.Msg = "获取记录成功"
	return resp
}
