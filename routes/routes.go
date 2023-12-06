// cmd/routes.go
// routes.go

package routes

import (
	"github.com/ZeinapIs/projectt/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// ... (other book-related routes)

	// Recipe-related routes
	app.Get("/api/recipes", handlers.GetAllRecipes)
	app.Get("/api/recipes/:recipeID", handlers.GetRecipeDetails)
	app.Post("/api/recipes", handlers.AddNewRecipe)
	app.Post("/api/recipes/:recipeID/mark-as-cooked", handlers.MarkAsCooked)
	app.Post("/api/recipes/:recipeID/mark-as-favorite", handlers.MarkAsFavorite)
	app.Get("/api/cooked-recipes", handlers.GetCookedRecipesList)
	app.Get("/api/favorite-recipes", handlers.GetFavoriteRecipesList)

	// ... (other book-related routes)
}
