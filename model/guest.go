package model

import (
	"github.com/jinzhu/gorm"
)

// Guest 房客模型
type Guest struct {
	gorm.Model
	RoomID   string `gorm:"unique"`
	IDNumber string
}

// GetGuest 用ID获取用户
func GetGuest(ID interface{}) (Guest, error) {
	var guest Guest
	result := DB.First(&guest, ID)
	return guest, result.Error
}

// CheckGuestIDNumber 检查身份证号
func (guest *Guest) CheckGuestIDNumber(idno string) bool {
	return idno == guest.IDNumber
}
