// cmd/main.go

package main

import (
	"github.com/ZeinapIs/projectt/database"
	"github.com/ZeinapIs/projectt/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Fiber app and database connection
	database.ConnectDb()
	app := fiber.New()
	routes.SetupRoutes(app)
	// Start the Fiber app
	app.Listen(":3000")
}
