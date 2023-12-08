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
	app.Post("/recipe", handlers.AddNewRecipe)
	app.Post("/api/recipes/:recipeID/mark-as-coi=oking", handlers.MarkAsCooking)
	app.Post("/api/recipes/:recipeID/mark-as-favorite", handlers.MarkAsToCook)
	app.Get("/api/cooked-recipes", handlers.GetCookingRecipesList)
	app.Get("/api/to-cook-recipes", handlers.GetToCookRecipesList)
	// MarkAsTried
	app.Post("/api/recipes/:recipeID/mark-as-tried", handlers.MarkAsTried)
	// Add the new route for fetching recipes by status
	// Update the route for fetching recipes by status

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
	app.Delete("/api/recipes/:recipeID", handlers.DeleteRecipe)

	app.Get("/recipe", handlers.NewRecipeView)
}
