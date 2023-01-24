package routes

import (
	"github.com/alvinscheibe/go-fiber-api/database"
	"github.com/alvinscheibe/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{
		ID:           productModel.ID,
		Name:         productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(context *fiber.Ctx) error {
	var product models.Product

	if err := context.BodyParser(&product); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)

	responseProduct := CreateResponseProduct(product)

	return context.Status(200).JSON(responseProduct)
}

func GetProducts(context *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)

	responseProducts := []Product{}
	for _, product := range products {
		responseProducts = append(responseProducts, CreateResponseProduct(product))
	}

	return context.Status(200).JSON(responseProducts)
}
