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
