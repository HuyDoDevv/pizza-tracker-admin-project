package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"pizza-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

type CustomerData struct {
	Title    string
	Order    models.Order
	Statuses []string
}

type OrderFormData struct {
	PizzaTypes []string
	PizzaSizes []string
}

type OrderRequest struct {
	Name         string   `form:"name" binding:"required,min=2,max=100"`
	Phone        string   `form:"phone" binding:"required,max=10"`
	Address      string   `form:"address" binding:"required,max=200"`
	PizzaSizes   []string `form:"pizza_size" binding:"required,min=1,dive,valid_pizza_size"`
	PizzaTypes   []string `form:"pizza_type" binding:"required,min=1,dive,valid_pizza_type"`
	Instructions []string `form:"instructions" binding:"max=200"`
}

func (h *Handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", OrderFormData{
		PizzaTypes: models.PizzaTypes,
		PizzaSizes: models.PizzaSizes,
	})
}

func (h *Handler) HandlerNewOrderPost(c *gin.Context) {
	var form OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItems := make([]models.OrderItem, len(form.PizzaSizes))
	for i := range orderItems {
		orderItems[i] = models.OrderItem{
			OrderID:     "fffff",
			PizzaSize:   form.PizzaSizes[i],
			PizzaType:   form.PizzaTypes[i],
			Instruction: form.Instructions[i],
		}
	}

	order := models.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       models.OrderStatus[0],
		Items:        orderItems,
	}

	fmt.Println("odhvduvhsiudhvisdvuisdvhsiduv")

	if err := h.orders.CreateOrder(&order); err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Something went wrong")
		return
	}

	slog.Info("Order created", "orderId", order.ID, "customer", order.CustomerName)

	c.Redirect(http.StatusSeeOther, "/customer/"+order.ID)
}

func (h *Handler) ServerCustomer(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.String(http.StatusBadRequest, "Order id is required")
		return
	}

	order, err := h.orders.GetOrderByID(orderID)
	if err != nil {
		c.String(http.StatusBadRequest, "Order not found")
		return
	}

	c.HTML(http.StatusOK, "customer.tmpl", CustomerData{
		Title:    "Pizza Order status" + orderID,
		Order:    *order,
		Statuses: models.OrderStatus,
	})
}
