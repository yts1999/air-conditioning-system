package middleware

import (
	"centralac/model"
	"centralac/serializer"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentGuest 获取登录房客
func CurrentGuest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		guestID := session.Get("guest_id")
		if guestID != nil {
			guest, err := model.GetGuest(guestID)
			if err == nil {
				c.Set("guest", &guest)
			}
		}
		c.Next()
	}
}

// GuestAuthRequired 需要登录
func GuestAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if guest, _ := c.Get("guest"); guest != nil {
			if _, ok := guest.(*model.Guest); ok {
				c.Next()
				return
			}
		}

		c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
