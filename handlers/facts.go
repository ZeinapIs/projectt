package handlers

import (
	"fmt"

	"github.com/ZeinapIs/projectt/database"
	"github.com/ZeinapIs/projectt/models"
	"github.com/gofiber/fiber/v2"
)

// GetAllRecipes retrieves a list of all recipes from the database
func GetAllRecipes(c *fiber.Ctx) error {
	fmt.Println("Fetching all recipes")
	var recipes []models.Recipe
	database.DB.Db.Find(&recipes)
	return c.Render("index", fiber.Map{
		"Title":    "My Recipe Collection",
		"Subtitle": "All of the recipes I own:",
		"Recipes":  recipes,
	})
}

// GetRecipeDetails retrieves details of a specific recipe by ID
func GetRecipeDetails(c *fiber.Ctx) error {
	recipeID := c.Params("recipeID")
	var recipe models.Recipe
	result := database.DB.Db.First(&recipe, recipeID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}
	return c.JSON(recipe)
}

// NewRecipeView renders the page to add a new recipe
func NewRecipeView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Recipe",
		"Subtitle": "Create a new recipe to add to your collection:",
	})
}

// AddNewRecipe adds a new recipe to the database
func AddNewRecipe(c *fiber.Ctx) error {
	var newRecipe models.Recipe
	if err := c.BodyParser(&newRecipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	database.DB.Db.Create(&newRecipe)
	return ConfirmationView(c, "Recipe added successfully", "You can add more recipes to your collection")
}

// UpdateRecipe updates an existing recipe's details
func UpdateRecipe(c *fiber.Ctx) error {
	recipeID := c.Params("recipeID")
	var updatedRecipe models.Recipe
	if err := c.BodyParser(&updatedRecipe); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	result := database.DB.Db.Model(&models.Recipe{}).Where("id = ?", recipeID).Updates(&updatedRecipe)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}
	return c.JSON(updatedRecipe)
}

// DeleteRecipe removes a recipe from the database
func DeleteRecipe(c *fiber.Ctx) error {
	recipeID := c.Params("recipeID")
	result := database.DB.Db.Delete(&models.Recipe{}, recipeID)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}
	return ConfirmationView(c, "Recipe deleted successfully", "You can manage your recipe collection")
}

// MarkAsCooking updates the status of a recipe to 'cooking'
func MarkAsCooking(c *fiber.Ctx) error {
	return markRecipeStatus(c, "cooking")
}

// MarkAsTried updates the status of a recipe to 'tried'
func MarkAsTried(c *fiber.Ctx) error {
	return markRecipeStatus(c, "tried")
}

// MarkAsNotTried updates the status of a recipe to 'not-tried'
func MarkAsNotTried(c *fiber.Ctx) error {
	return markRecipeStatus(c, "not-tried")
}

// MarkAsToCook updates the status of a recipe to 'to-cook'
func MarkAsToCook(c *fiber.Ctx) error {
	return markRecipeStatus(c, "to-cook")
}

// markRecipeStatus updates the status of a recipe
func markRecipeStatus(c *fiber.Ctx, status string) error {
	recipeID := c.Params("recipeID")
	var recipe models.Recipe
	result := database.DB.Db.First(&recipe, recipeID)
	if result.Error != nil {
		fmt.Println("Error fetching recipe:", result.Error)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}
	recipe.Status = status
	database.DB.Db.Save(&recipe)
	return c.JSON(recipe)
}

// ConfirmationView renders a confirmation view with a custom message
func ConfirmationView(c *fiber.Ctx, title, subtitle string) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    title,
		"Subtitle": subtitle,
	})
}

// SearchRecipesByIngredient searches for recipes based on a partial match of ingredients
func SearchRecipesByIngredient(c *fiber.Ctx) error {
	partialIngredient := c.Params("partialIngredient")
	var recipes []models.Recipe
	err := database.DB.Db.Where("Ingredients ILIKE ?", "%"+partialIngredient+"%").Find(&recipes).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(recipes)
}

// SearchRecipesByInstruction searches for recipes based on a partial match of instructions
func SearchRecipesByInstruction(c *fiber.Ctx) error {
	partialInstruction := c.Params("partialInstruction")
	var recipes []models.Recipe
	err := database.DB.Db.Where("Instructions ILIKE ?", "%"+partialInstruction+"%").Find(&recipes).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(recipes)
}

// SearchRecipesByTitle searches for recipes based on a partial match of the title
func SearchRecipesByTitle(c *fiber.Ctx) error {
	partialTitle := c.Params("partialTitle")
	var recipes []models.Recipe
	err := database.DB.Db.Where("Title ILIKE ?", "%"+partialTitle+"%").Find(&recipes).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.JSON(recipes)

}
func GetCookingRecipes(c *fiber.Ctx) error {
	return getRecipesByStatusHelper(c, "cooking")
}

func GetToCookRecipes(c *fiber.Ctx) error {
	return getRecipesByStatusHelper(c, "to-cook")
}

func GetTriedRecipes(c *fiber.Ctx) error {
	return getRecipesByStatusHelper(c, "tried")
}

func GetNotTriedRecipes(c *fiber.Ctx) error {
	return getRecipesByStatusHelper(c, "not-tried")
}

func getRecipesByStatusHelper(c *fiber.Ctx, status string) error {
	var recipes []models.Recipe
	err := database.DB.Db.Where("status = ?", status).Find(&recipes).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(recipes)
}

func SearchRecipes(c *fiber.Ctx) error {
	// Get the search term from the query string
	searchTerm := c.Query("query")

	// Construct the query for a partial match
	query := "%" + searchTerm + "%"

	// Fetch recipes from the database with a partial match
	var searchResults []models.Recipe
	result := database.DB.Db.Where("Title ILIKE ? OR Ingredients ILIKE ? OR Instructions ILIKE ?", query, query, query).Find(&searchResults)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(searchResults)
}

func GetRecipesByStatus(c *fiber.Ctx) error {
	status := c.Params("status")

	var recipes []models.Recipe
	err := database.DB.Db.Where("status = ?", status).Find(&recipes).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.JSON(recipes)
}
