package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func setupRouters(router *gin.Engine, h *Handler, store sessions.Store) {
	router.Use(sessions.Sessions("pizza-tracker", store))

	router.GET("/", h.ServeNewOrderForm)
	router.POST("/new-order", h.HandlerNewOrderPost)
	router.GET("/customer/:id", h.ServerCustomer)

	router.GET("/login", h.HandleLoginGet)
	router.POST("/login", h.HandleLoginPost)
	router.POST("/logout", h.HandleLogout)

	admin := router.Group("/admin")
	admin.Use(h.AuthMiddleware())
	{
		admin.GET("", h.ServerAdminDashboard)
		admin.POST("/order/:id/update", h.HandlerOrderPut)
		admin.POST("/order/:id/delete", h.HandlerOrderDelete)
	}
	router.Static("/static", "./templates/static")
}
