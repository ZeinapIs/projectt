// cmd/routes.go
// routes.go

package routes

import (
	"github.com/ZeinapIs/projectt/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlers.GetAllRecipes)
	app.Get("/api/recipes/:recipeID", handlers.GetRecipeDetails)
	app.Post("/api/recipes", handlers.AddNewRecipe)
	app.Post("/api/recipes/:recipeID/mark-as-cooking", handlers.MarkAsCooking)
	app.Post("/api/recipes/:recipeID/mark-as-cook", handlers.MarkAsToCook)
	app.Post("/api/recipes/:recipeID/mark-as-tried", handlers.MarkAsTried)
	app.Get("/api/recipes/status/cooking", handlers.GetRecipesByStatus)
	app.Get("/api/recipes/status/to-cook", handlers.GetRecipesByStatus)
	app.Get("/api/recipes/status/tried", handlers.GetRecipesByStatus)
	app.Get("/api/recipes/status/not-tried", handlers.GetRecipesByStatus)

	app.Post("/api/recipes/:recipeID/mark-as-not-tried", handlers.MarkAsNotTried)

	app.Get("/api/recipes/ingr/:partialIngredient", handlers.SearchRecipesByIngredient)

	app.Get("/api/recipes/instr/:partialInstruction", handlers.SearchRecipesByInstruction)

	app.Get("/api/recipes/title/:partialTitle", handlers.SearchRecipesByTitle)
	app.Get("/api/recipes/search/:searchTerm", handlers.SearchRecipes)

	app.Put("/api/recipes/:recipeID", handlers.UpdateRecipe)
	app.Patch("/api/recipes/:recipeID", handlers.UpdateRecipe)
	app.Delete("/api/recipes/:recipeID", handlers.DeleteRecipe)
	app.Get("/api/recipes/status/:status", handlers.GetRecipesByStatus)

	app.Get("/recipe", handlers.NewRecipeView)
}
