package api

import (
	"centralac/service"

	"github.com/gin-gonic/gin"
)

// RecordOfRoomList 获取房间温控记录接口
func RecordOfRoomList(c *gin.Context) {
	var service service.RecordOfRoomService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
