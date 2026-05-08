package main

import "pizza-tracker/internal/models"

type Handler struct {
	orders       *models.OrderModel
	users        *models.UserModel
	notification *NotificationManager
}

func NewHandler(dbModel *models.DBModel) *Handler {
	return &Handler{
		orders:       &dbModel.Order,
		users:        &dbModel.User,
		notification: NewNotificationManager(),
	}
}
