package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouters(router *gin.Engine, h *Handler) {
	router.GET("/", h.ServeNewOrderForm)
	router.POST("/new-order", h.HandlerNewOrderPost)
	router.GET("/customer/:id", h.ServerCustomer)

	router.Static("/static", "/templates/static")
}
