package routes

import (
	"errors"
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

func GetUsers(context *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)

	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return context.Status(200).JSON(responseUsers)
}

func GetUser(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var user models.User

	if err != nil {
		return context.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)

	return context.Status(200).JSON(responseUser)
}

func UpdateUser(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")

	var user models.User

	if err != nil {
		return context.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findUser(id, &user); err != nil {
		return context.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := context.BodyParser(&updateData); err != nil {
		return context.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseUser(user)

	return context.Status(200).JSON(responseUser)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)

	if user.ID == 0 {
		return errors.New("User doesn't exist")
	}

	return nil
}