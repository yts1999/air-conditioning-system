package server

import (
	"centralac/api"
	"centralac/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 房客登录
		v1.POST("guest/login", api.GuestLogin)
		// 从控机开关机
		v1.POST("room/startup", api.RoomStartup)
		v1.POST("room/shutdown", api.RoomShutdown)
		// 管理员注册、登录
		v1.POST("admin/register", api.AdminRegister)
		v1.POST("admin/login", api.AdminLogin)

		// 需要房客登录保护的
		guestAuth := v1.Group("")
		{
			guestAuth.Use(middleware.CurrentGuest())
			guestAuth.Use(middleware.GuestAuthRequired())

			guestAuth.GET("guest/me", api.GuestMe)
			guestAuth.DELETE("guest/logout", api.GuestLogout)

			guestAuth.POST("room/startup", api.RoomStartup)
			guestAuth.DELETE("room/shutdown", api.RoomShutdown)

			guestAuth.POST("room/updateCurrentTemp", api.RoomCurrentTempUpdate)
			guestAuth.POST("room/updateTargetTemp", api.RoomTargetTempUpdate)
			guestAuth.POST("room/updateWindSpeed", api.RoomWindSpeedUpdate)

			guestAuth.POST("wind/start", api.RoomWindStart)
			guestAuth.POST("wind/stop", api.RoomWindStop)
		}

		// 需要管理员登录保护的
		adminAuth := v1.Group("")
		{
			adminAuth.Use(middleware.CurrentAdmin())
			adminAuth.Use(middleware.AdminAuthRequired())

			adminAuth.DELETE("admin/logout", api.AdminLogout)

			adminAuth.POST("guest/register", api.GuestRegister)
			adminAuth.DELETE("guest/delete", api.GuestDelete)

			adminAuth.POST("center/startup", api.CenterStartup)
			adminAuth.DELETE("center/shutdown", api.CenterShutdown)
			adminAuth.POST("center/change", api.CenterChangeWorkMode)

			adminAuth.POST("room/create", api.RoomCreate)
			adminAuth.DELETE("room/delete", api.RoomDelete)
			adminAuth.GET("room/show", api.RoomShow)
			adminAuth.GET("room/list", api.RoomList)

			adminAuth.GET("record/all", api.RecordOfRoomList)
			adminAuth.GET("record/day", api.RecordOfRoomOfDayList)
			adminAuth.GET("record/month", api.RecordOfRoomOfMonthList)
			adminAuth.GET("record/year", api.RecordOfRoomOfYearList)
		}
	}
	return r
}
