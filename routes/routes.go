// cmd/routes.go
// routes.go

package routes

import (
	"github.com/ZeinapIs/projectt/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", handlers.GetAllRecipes)
	app.Get("/recipe/:recipeID", handlers.GetRecipeDetails)
	app.Post("/recipe", handlers.AddNewRecipe)
	app.Post("/api/recipes/:recipeID/cooking", handlers.MarkAsCooking)
	app.Post("/api/recipes/:recipeID/cook", handlers.MarkAsToCook)
	app.Post("/api/recipes/:recipeID/tried", handlers.MarkAsTried)
	app.Get("/cooking", handlers.GetCookingRecipes)
	app.Get("/to-cook", handlers.GetToCookRecipes)
	app.Get("/tried", handlers.GetTriedRecipes)
	app.Get("/not-tried", handlers.GetNotTriedRecipes)

	app.Post("/api/recipes/:recipeID/not-tried", handlers.MarkAsNotTried)

	app.Get("/api/recipes/ingr/:partialIngredient", handlers.SearchRecipesByIngredient)

	app.Get("/api/recipes/instr/:partialInstruction", handlers.SearchRecipesByInstruction)

	app.Get("/api/recipes/title/:partialTitle", handlers.SearchRecipesByTitle)
	app.Get("/api/recipes/search/:searchTerm", handlers.SearchRecipes)
	app.Get("/recipe/:recipeID/edit", handlers.EditRecipe)

	app.Patch("/recipe/:recipeID", handlers.UpdateRecipe)
	app.Delete("recipe/delete/:recipeID", handlers.DeleteRecipe)

	app.Get("/recipe", handlers.NewRecipeView)
	app.Get("/api/recipes", handlers.GetAllRecipesAsJSON)

}
