package api

import (
	"centralac/service"

	"github.com/gin-gonic/gin"
)

// CenterStartup 中央空调开机接口
func CenterStartup(c *gin.Context) {
	var service service.CenterStartupService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Startup()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CenterShutdown 中央空调关机接口
func CenterShutdown(c *gin.Context) {
	var service service.CenterShutdownService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Shutdown()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CenterChangeWorkMode 中央空调改变工作模式接口
func CenterChangeWorkMode(c *gin.Context) {
	var service service.CenterChangeWorkModeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Change()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// CenterStatus 获取中央空调状态接口
func CenterStatus(c *gin.Context) {
	var service service.CenterStatusService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Status()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
