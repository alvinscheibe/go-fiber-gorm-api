package routes

import (
	"errors"
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

func GetProduct(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var product models.Product

	if err != nil {
		return context.Status(400).JSON("Please ensure that :id as an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	responseProduct := CreateResponseProduct(product)

	return context.Status(200).JSON(responseProduct)
}

func UpdateProduct(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var product models.Product

	if err != nil {
		return context.Status(400).JSON("Please ensure that :id as an integer")
	}

	if err := findProduct(id, &product); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateProduct UpdateProduct

	if err := context.BodyParser(updateProduct); err != nil {
		return context.Status(500).JSON(err.Error())
	}

	product.Name = updateProduct.Name
	product.SerialNumber = updateProduct.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return context.Status(200).JSON(responseProduct)
}

func findProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("Product doesn't exist")
	}

	return nil
}
