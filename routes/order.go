package routes

import (
	"github.com/alvinscheibe/go-fiber-api/models"
	"time"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateOrderResponse(orderModel models.Order, user User, product Product) Order {
	return Order{
		ID:        orderModel.ID,
		User:      user,
		Product:   product,
		CreatedAt: orderModel.CreatedAt,
	}
}
