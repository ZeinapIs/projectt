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
	app.Get("/recipes/cooking", handlers.GetCookingRecipesList)
	app.Get("/recipes/to-cook", handlers.GetToCookRecipesList)
	app.Post("/api/recipes/:recipeID/mark-as-tried", handlers.MarkAsTried)
	app.Get("/recipes/tried", handlers.GetTriedRecipesList)
	app.Get("/recipes/not-tried", handlers.GetNotTriedRecipesList)
	app.Post("/api/recipes/:recipeID/mark-as-not-tried", handlers.MarkAsNotTried)

	app.Get("/api/recipes/ingr/:partialIngredient", handlers.SearchRecipesByIngredients)

	app.Get("/api/recipes/instr/:partialInstruction", handlers.SearchRecipesByInstructions)

	app.Get("/api/recipes/title/:partialTitle", handlers.SearchRecipesByTitle)
	app.Put("/api/recipes/:recipeID", handlers.UpdateRecipe)
	app.Patch("/api/recipes/:recipeID", handlers.UpdateRecipe)
	app.Delete("/api/recipes/:recipeID", handlers.DeleteRecipe)

	app.Get("/recipe", handlers.NewRecipeView)
}
