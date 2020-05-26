package api

import (
	"centralac/service"

	"github.com/gin-gonic/gin"
)

// RecordOfRoomList 获取房间所有温控记录接口
func RecordOfRoomList(c *gin.Context) {
	var service service.RecordOfRoomService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RecordOfRoomOfDayList 获取房间日温控记录接口
func RecordOfRoomOfDayList(c *gin.Context) {
	var service service.RecordOfRoomOfDayService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RecordOfRoomOfMonthList 获取房间月温控记录接口
func RecordOfRoomOfMonthList(c *gin.Context) {
	var service service.RecordOfRoomOfMonthService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RecordOfRoomOfYearList 获取房间年温控记录接口
func RecordOfRoomOfYearList(c *gin.Context) {
	var service service.RecordOfRoomOfYearService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
