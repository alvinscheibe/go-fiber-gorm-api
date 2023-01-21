package routes

import (
	"github.com/alvinscheibe/go-fiber-api/database"
	"github.com/alvinscheibe/go-fiber-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(context *fiber.Ctx) error {
	var user models.User

	if err := context.BodyParser(&user); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)

	return context.Status(200).JSON(responseUser)
}
