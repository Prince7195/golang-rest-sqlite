package routes

import (
	"errors"

	"github.com/Prince7195/golang-rest-sqlite/database"
	"github.com/Prince7195/golang-rest-sqlite/models"
	"github.com/gofiber/fiber/v2"
)

// https://gorm.io/docs/#Install

type User struct {
	// this is not a User model, see this as a serializer
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) User {
	return User{
		ID: userModel.ID,
		FirstName: userModel.FirstName,
		LastName: userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.DB.Create(&user)
	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.DB.Find(&users)

	responseUsers := []User{}

	for _, user := range users {
		resUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, resUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")
	}

	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an intiger")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an intiger")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.DB.Save(&user)

	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an intiger")
	}

	database.Database.DB.Delete(&user, id)

	response := CreateResponseUser(user)

	return c.Status(200).JSON(response)
}
