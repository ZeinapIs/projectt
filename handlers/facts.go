// handlers/recipes.go

package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/ZeinapIs/projectt/database"
	"github.com/ZeinapIs/projectt/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllRecipes retrieves a list of all recipes
func GetAllRecipes(c *fiber.Ctx) error {
	// Implement the logic to fetch all recipes from the database
	var recipes []models.Recipe
	database.DB.Db.Find(&recipes)
	return c.JSON(recipes)
}

// GetRecipeDetails retrieves details of a specific recipe
func GetRecipeDetails(c *fiber.Ctx) error {
	// Implement the logic to fetch details of a specific recipe from the database
	recipeID := c.Params("recipeID")
	var recipe models.Recipe
	result := database.DB.Db.First(&recipe, recipeID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}
	return c.JSON(recipe)
}

// AddNewRecipe adds a new recipe to the database
func AddNewRecipe(c *fiber.Ctx) error {
	var newRecipe models.Recipe

	// Read the request body
	bodyBytes := c.Body()

	// Unmarshal the request body into the Recipe struct
	if err := json.Unmarshal(bodyBytes, &newRecipe); err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Insert the new recipe into the database
	database.DB.Db.Create(&newRecipe)

	return c.JSON(newRecipe)
}

// MarkAsCooked marks a recipe as cooked
func MarkAsCooked(c *fiber.Ctx) error {
	return markRecipeStatus(c, "cooked")
}

// MarkAsFavorite marks a recipe as a favorite
func MarkAsFavorite(c *fiber.Ctx) error {
	return markRecipeStatus(c, "favorite")
}

// GetCookedRecipesList retrieves a list of cooked recipes
func GetCookedRecipesList(c *fiber.Ctx) error {
	return getRecipesByStatus(c, "cooked")
}

// GetFavoriteRecipesList retrieves a list of favorite recipes
func GetFavoriteRecipesList(c *fiber.Ctx) error {
	return getRecipesByStatus(c, "favorite")
}

// ... (other recipe-related handlers)

func markRecipeStatus(c *fiber.Ctx, status string) error {
	// Implement the logic to update the recipe status (similar to markBookStatus)
	return c.SendStatus(fiber.StatusNoContent)
}

func getRecipesByStatus(c *fiber.Ctx, status string) error {
	// Implement the logic to fetch recipes by status from the database (similar to getBooksByStatus)
	var recipes []models.Recipe
	database.DB.Db.Where("status = ?", status).Find(&recipes)
	return c.JSON(recipes)
}
