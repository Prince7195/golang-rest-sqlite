package main

import (
	"log"

	"github.com/Prince7195/golang-rest-sqlite/database"
	"github.com/Prince7195/golang-rest-sqlite/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to Golang Fiber API")
}

func setupRoutes(app *fiber.App) {
	// WELCOME API
	app.Get("/api", welcome)

	// USER ENDPOINTS
	// CREATE A USER
	app.Post("/api/user", routes.CreateUser)

	// GET ALL USERS
	app.Get("/api/users", routes.GetUsers)

	// GET USER BY ID
	app.Get("/api/user/:id", routes.GetUser)

	// UPDATE USER BY ID
	app.Put("/api/user/:id", routes.UpdateUser)

	// DELETE USER BY ID
	app.Delete("/api/user/:id", routes.DeleteUser)

	// PRODUCT ENDPOINTS
	// CREATE A PRODUCT
	app.Post("/api/product", routes.CreateProduct)

	// GET ALL PRODUCTS
	app.Get("/api/products", routes.GetProducts)

	// GET PRODUCT BY ID
	app.Get("/api/product/:id", routes.GetProduct)

	// UPDATE PRODUCT BY ID
	app.Put("/api/product/:id", routes.UpdateProduct)

	// DELETE PRODUCT BY ID
	app.Delete("/api/product/:id", routes.DeleteProduct)

	// ORDER ENDPOINTS
	// CREATE AN ORDER
	app.Post("/api/order", routes.CreateOrder)

	// GET ALL ORDERS
	app.Get("/api/orders", routes.GetOrders)

	// GET ORDER BY ID
	app.Get("/api/order/:id", routes.GetOrder)
}

func main() {
	database.ConnectDB()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":8000"))
}
