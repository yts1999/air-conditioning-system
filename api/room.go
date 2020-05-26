package api

import (
	"centralac/service"

	"github.com/gin-gonic/gin"
)

// RoomCreate 创建房间接口
func RoomCreate(c *gin.Context) {
	var service service.RoomCreateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomDelete 删除房间接口
func RoomDelete(c *gin.Context) {
	var service service.RoomDeleteService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomShow 获取房间信息接口
func RoomShow(c *gin.Context) {
	var service service.RoomShowService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomList 获取所有房间信息接口
func RoomList(c *gin.Context) {
	var service service.RoomListService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
