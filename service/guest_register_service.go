package service

import (
	"centralac/model"
	"centralac/serializer"
	"regexp"
	"strconv"
)

var (
	wi      = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	vercode = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
)

// GuestRegisterService 管理房客注册服务
type GuestRegisterService struct {
	RoomID   string `form:"room_id" json:"room_id" binding:"required,min=3,max=4"`
	IDNumber string `form:"id_number" json:"id_number" binding:"required,min=18,max=18"`
}

// valid 验证表单
func (service *GuestRegisterService) valid() *serializer.Response {
	count := 0
	model.DB.Model(&model.Guest{}).Where("room_id = ?", service.RoomID).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "房间已经注册",
		}
	}
	pattern := `^\d{17}[\dX]$`
	result, _ := regexp.MatchString(pattern, service.IDNumber)
	if !result {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "身份证号错误",
		}
	}
	sum := 0
	for i, ch := range service.IDNumber[:len(service.IDNumber)-1] {
		chnum, _ := strconv.Atoi(string(ch))
		sum += chnum * wi[i]
	}
	if vercode[sum%11] != service.IDNumber[len(service.IDNumber)-1] {
		return &serializer.Response{
			Code: serializer.CodeParamErr,
			Msg:  "身份证号错误",
		}
	}
	return nil
}

// Register 用户注册
func (service *GuestRegisterService) Register() serializer.Response {
	guest := model.Guest{
		RoomID:   service.RoomID,
		IDNumber: service.IDNumber,
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 创建用户
	if err := model.DB.Create(&guest).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	resp := serializer.BuildGuestResponse(guest)
	resp.Msg = "注册成功"
	return resp
}
