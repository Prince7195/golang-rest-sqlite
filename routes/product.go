package routes

import (
	"errors"

	"github.com/Prince7195/golang-rest-sqlite/database"
	"github.com/Prince7195/golang-rest-sqlite/models"
	"github.com/gofiber/fiber/v2"
)

// https://gorm.io/docs/#Install

type Product struct {
	// this is not a User model, see this as a serializer
	ID uint `json:"id"`
	Name string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{
		ID: productModel.ID,
		Name: productModel.Name,
		SerialNumber: productModel.SerialNumber,
	}
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	database.Database.DB.Create(&product)
	response := CreateResponseProduct(product)

	return c.Status(200).JSON(response)
}

func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.DB.Find(&products)

	responseProducts := []Product{}

	for _, product := range products {
		resProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, resProduct)
	}

	return c.Status(200).JSON(responseProducts)
}

func findProduct(id int, product *models.Product) error {
	database.Database.DB.Find(&product, "id = ?", id)
	if product.ID == 0 {
		return errors.New("Product does not exist")
	}

	return nil
}

func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an intiger")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := CreateResponseProduct(product)

	return c.Status(200).JSON(response)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an intiger")
	}

	if err := findProduct(id, &product); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateProduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.DB.Save(&product)

	response := CreateResponseProduct(product)

	return c.Status(200).JSON(response)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an intiger")
	}

	database.Database.DB.Delete(&product, id)

	response := CreateResponseProduct(product)

	return c.Status(200).JSON(response)
}
