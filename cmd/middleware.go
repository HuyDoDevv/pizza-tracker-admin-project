package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := GetSessionString(c, "userId")
		if userId == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		_, err := h.users.GetUserById(userId)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
