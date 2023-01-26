package routes

import (
	"errors"
	"github.com/alvinscheibe/go-fiber-api/database"
	"github.com/alvinscheibe/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateResponseOrder(orderModel models.Order, user User, product Product) Order {
	return Order{
		ID:        orderModel.ID,
		User:      user,
		Product:   product,
		CreatedAt: orderModel.CreatedAt,
	}
}

func CreateOrder(context *fiber.Ctx) error {
	var order models.Order

	if err := context.BodyParser(&order); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return context.Status(200).JSON(responseOrder)
}

func GetOrders(context *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product

		if err := findUser(order.UserRefer, &user); err != nil {
			return context.Status(500).JSON(err.Error())
		}

		if err := findProduct(order.ProductRefer, &product); err != nil {
			return context.Status(500).JSON(err.Error())
		}

		responseUser := CreateResponseUser(user)
		responseProduct := CreateResponseProduct(product)
		responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

		responseOrders = append(responseOrders, responseOrder)
	}
	
	return context.Status(200).JSON(responseOrders)
}

func findOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)

	if order.ID == 0 {
		return errors.New("Order doesn't exist")
	}

	return nil
}