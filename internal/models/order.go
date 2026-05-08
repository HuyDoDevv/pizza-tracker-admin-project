package models

import (
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

var (
	OrderStatus = []string{
		"Order placed",
		"Preparing",
		"Baking",
		"Quality check",
		"Ready for delivery",
	}
	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"Hawaiian",
		"Veggie",
		"BBQ Chicken",
	}

	PizzaSizes = []string{
		"Small",
		"Medium",
		"Large",
		"X Large",
	}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID           string      `json:"id" gorm:"primaryKey;size:14"`
	Status       string      `json:"status" gorm:"not null"`
	CustomerName string      `json:"customer_name" gorm:"not null"`
	Phone        string      `json:"phone" gorm:"not null"`
	Address      string      `json:"address" gorm:"not null"`
	Items        []OrderItem `json:"pizzas" gorm:"foreignKey:OrderID"`
	CreatedAt    string      `json:"created_at"`
}

type OrderItem struct {
	ID          string `json:"id" gorm:"primaryKey;size:14"`
	OrderID     string `json:"order_id" gorm:"index;not null;size:14"`
	PizzaType   string `json:"pizza_type" gorm:"not null"`
	PizzaSize   string `json:"pizza_size" gorm:"not null"`
	Instruction string `json:"instruction"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	if o.ID == "" {
		o.ID = shortid.MustGenerate()
	}
	return nil
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) error {
	if oi.ID == "" {
		oi.ID = shortid.MustGenerate()
	}
	return nil
}

func (m *OrderModel) CreateOrder(order *Order) error {
	return m.DB.Create(order).Error
}

func (m *OrderModel) GetOrderByID(id string) (*Order, error) {
	var order Order
	err := m.DB.Preload("Items").First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderModel) GetAllOrders() ([]Order, error) {
	var orders []Order
	err := o.DB.Preload("Items").Order("created_at desc").Find(&orders).Error
	return orders, err
}

func (o *OrderModel) UpdateOrderStatus(id string, status string) error {
	return o.DB.Model(&Order{}).Where("id = ?", id).Update("status", status).Error
}

func (o *OrderModel) DeleteOrder(id string) error {
	return o.DB.Select("Items").Delete(&Order{ID: id}).Error
}
