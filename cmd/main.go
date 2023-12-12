// cmd/main.go

package main

import (
	"github.com/ZeinapIs/projectt/database"
	"github.com/ZeinapIs/projectt/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize Fiber app and database connection
	database.ConnectDb()
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	routes.SetupRoutes(app)

	app.Static("/", "./public")

	app.Listen(":3000")
}
