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

// MarkAsTried marks a recipe as tried
func MarkAsTried(c *fiber.Ctx) error {
	return markRecipeStatus(c, "tried")
}

// GetTriedRecipesList retrieves a list of tried recipes
func GetTriedRecipesList(c *fiber.Ctx) error {
	return getRecipesByStatus(c, "tried")
}

// MarkAsNotTried marks a recipe as not tried
func MarkAsNotTried(c *fiber.Ctx) error {
	return markRecipeStatus(c, "not tried")
}

// GetNotTriedRecipesList retrieves a list of not tried recipes
func GetNotTriedRecipesList(c *fiber.Ctx) error {
	return getRecipesByStatus(c, "not tried")
}

// SearchRecipesByIngredients searches for recipes based on ingredients
// SearchRecipesByIngredients searches for recipes based on a partial match of ingredients
func SearchRecipesByIngredients(c *fiber.Ctx) error {
	// Get the partial ingredient from the URL parameter
	partialIngredient := c.Params("partialIngredient")

	// Construct the query for a partial match
	query := "%" + partialIngredient + "%"

	// Fetch recipes from the database with a partial ingredient match
	var searchResults []models.Recipe
	result := database.DB.Db.Where("Ingredients ILIKE ?", query).Find(&searchResults)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(searchResults)
}

// SearchRecipesByInstructions searches for recipes based on a partial match of instructions
func SearchRecipesByInstructions(c *fiber.Ctx) error {
	// Get the partial instruction from the URL parameter
	partialInstruction := c.Params("partialInstruction")

	// Construct the query for a partial match
	query := "%" + partialInstruction + "%"

	// Fetch recipes from the database with a partial instruction match
	var searchResults []models.Recipe
	result := database.DB.Db.Where("Instructions ILIKE ?", query).Find(&searchResults)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(searchResults)
}

// SearchRecipesByTitle searches for recipes based on a partial match of the title
func SearchRecipesByTitle(c *fiber.Ctx) error {
	// Get the partial title from the URL parameter
	partialTitle := c.Params("partialTitle")

	// Construct the query for a partial match
	query := "%" + partialTitle + "%"

	// Fetch recipes from the database with a partial title match
	var searchResults []models.Recipe
	result := database.DB.Db.Where("Title ILIKE ?", query).Find(&searchResults)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(searchResults)
}
func UpdateRecipe(c *fiber.Ctx) error {
	// Get recipeID from params
	recipeID := c.Params("recipeID")

	// Retrieve the recipe from the database
	var existingRecipe models.Recipe
	result := database.DB.Db.First(&existingRecipe, recipeID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Read the request body
	bodyBytes := c.Body()

	// Unmarshal the request body into the Recipe struct
	var updatedRecipe models.Recipe
	if err := json.Unmarshal(bodyBytes, &updatedRecipe); err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Update the existing recipe with the new details
	existingRecipe.Title = updatedRecipe.Title
	existingRecipe.Ingredients = updatedRecipe.Ingredients
	existingRecipe.Instructions = updatedRecipe.Instructions

	// Save the updated recipe to the database
	database.DB.Db.Save(&existingRecipe)

	// Return the updated recipe
	return c.JSON(existingRecipe)
}
