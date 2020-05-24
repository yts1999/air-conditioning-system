package service

import (
	"centralac/model"
	"centralac/serializer"
)

// GuestDeleteService 管理删除房客的服务
type GuestDeleteService struct {
	RoomID string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
}

// Delete 删除房客函数
func (service *GuestDeleteService) Delete() serializer.Response {
	//检查房客是否存在
	var guest model.Guest
	if !model.DB.Where("room_id = ?", service.RoomID).First(&guest).RecordNotFound() {
		return serializer.ParamErr("房客不存在", nil)
	}

	// 删除房客
	if err := model.DB.Where("room_id = ?", service.RoomID).Delete(&guest).Error; err != nil {
		return serializer.ParamErr("房客删除失败", err)
	}

	resp := serializer.BuildGuestResponse(guest)
	resp.Msg = "房客删除成功"
	return resp
}
