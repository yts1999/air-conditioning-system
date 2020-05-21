package middleware

import (
	"centralac/model"
	"centralac/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentAdmin 获取登录管理员
func CurrentAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		adminID := session.Get("admin_id")
		if adminID != nil {
			admin, err := model.GetAdmin(adminID)
			if err == nil {
				c.Set("admin", &admin)
			}
		}
		c.Next()
	}
}

// AdminAuthRequired 需要登录
func AdminAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if admin, _ := c.Get("admin"); admin != nil {
			if _, ok := admin.(*model.Admin); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
