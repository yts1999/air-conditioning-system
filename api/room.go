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

// RoomStartup 从控机开机接口
func RoomStartup(c *gin.Context) {
	var service service.RoomStartupService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Startup()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomShutdown 从控机关机接口
func RoomShutdown(c *gin.Context) {
	var service service.RoomShutdownService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Shutdown()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomCurrentTempUpdate 更新当前温度接口
func RoomCurrentTempUpdate(c *gin.Context) {
	var service service.RoomCurrentTempUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomTargetTempUpdate 更改目标温度接口
func RoomTargetTempUpdate(c *gin.Context) {
	var service service.RoomTargetTempUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomWindSpeedUpdate 更改风速接口
func RoomWindSpeedUpdate(c *gin.Context) {
	var service service.RoomWindSpeedUpdateService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomWindStart 请求送风接口
func RoomWindStart(c *gin.Context) {
	var service service.RoomWindStartService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Start()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// RoomWindStop 停止送风接口
func RoomWindStop(c *gin.Context) {
	var service service.RoomWindStopService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Stop()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
