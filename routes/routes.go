// cmd/routes.go
// routes.go

package routes

import (
	"github.com/ZeinapIs/projectt/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/api/recipes", handlers.GetAllRecipes)
	app.Get("/api/recipes/:recipeID", handlers.GetRecipeDetails)
	app.Post("/api/recipes", handlers.AddNewRecipe)
	app.Post("/api/recipes/:recipeID/mark-as-cooked", handlers.MarkAsCooked)
	app.Post("/api/recipes/:recipeID/mark-as-favorite", handlers.MarkAsFavorite)
	app.Get("/api/cooked-recipes", handlers.GetCookedRecipesList)
	app.Get("/api/favorite-recipes", handlers.GetFavoriteRecipesList)
	// MarkAsTried
	app.Post("/api/recipes/:recipeID/mark-as-tried", handlers.MarkAsTried)

	// GetTriedRecipesList
	app.Get("/api/tried-recipes", handlers.GetTriedRecipesList)

	// MarkAsNotTried
	app.Post("/api/recipes/:recipeID/mark-as-not-tried", handlers.MarkAsNotTried)

	// GetNotTriedRecipesList
	app.Get("/api/not-tried-recipes", handlers.GetNotTriedRecipesList)

	// Search by ingredients
	app.Get("/api/recipes/ingr/:partialIngredient", handlers.SearchRecipesByIngredients)

	// Search by instructions
	app.Get("/api/recipes/instr/:partialInstruction", handlers.SearchRecipesByInstructions)

	app.Get("/api/recipes/title/:partialTitle", handlers.SearchRecipesByTitle)
	app.Put("/api/recipes/:recipeID", handlers.UpdateRecipe)
	app.Patch("/api/recipes/:recipeID", handlers.UpdateRecipe)

}
