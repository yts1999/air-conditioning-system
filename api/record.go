package api

import (
	"centralac/service"

	"github.com/gin-gonic/gin"
)

// RecordList 获取温控记录接口
func RecordList(c *gin.Context) {
	var service service.RecordService
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
