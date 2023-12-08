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
	return c.Render("index", fiber.Map{
		"Title":    "My Recipe Collection",
		"Subtitle": "All of the recipes I own:",
		"Recipes":  recipes,
	})
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
func NewRecipeView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"Title":    "New Recipe",
		"Subtitle": "Create a new recipe to add to your collection:",
	})
}

// AddNewRecipe adds a new recipe to the database
// AddNewRecipe adds a new recipe to the database
func AddNewRecipe(c *fiber.Ctx) error {
	var newRecipe models.Recipe

	// Read the request body
	bodyBytes := c.Body()

	// Unmarshal the request body into the Recipe struct
	if err := json.Unmarshal(bodyBytes, &newRecipe); err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload", "details": err.Error()})
	}

	// Insert the new recipe into the database
	database.DB.Db.Create(&newRecipe)

	// Customize the confirmation message for recipe addition
	title := "Recipe added successfully"
	subtitle := "You can add more recipes to your collection"
	confirmation := ConfirmationView(c, title, subtitle)

	// Return the created recipe details along with the confirmation message
	return c.JSON(fiber.Map{
		"confirmation": confirmation,
		"recipe":       newRecipe,
	})
}

// Assume ConfirmationView is the same for both books and recipes
func ConfirmationView(c *fiber.Ctx, title, subtitle string) error {
	return c.Render("confirmation", fiber.Map{
		"Title":    title,
		"Subtitle": subtitle,
	})
}

// MarkAsCooked marks a recipe as cooked
func MarkAsCooking(c *fiber.Ctx) error {
	return markRecipeStatus(c, "cooking")
}

func MarkAsTried(c *fiber.Ctx) error {
	return markRecipeStatus(c, "tried")
}

func MarkAsNotTried(c *fiber.Ctx) error {
	return markRecipeStatus(c, "not-tried")
}

func MarkAsToCook(c *fiber.Ctx) error {
	return markRecipeStatus(c, "to-cook")
}

// GetCookedRecipesList retrieves a list of cooked recipes
func GetCookingRecipesList(c *fiber.Ctx) error {
	return GetRecipesByStatus(c, "cooking")
}

// GetFavoriteRecipesList retrieves a list of favorite recipes
// GetTriedRecipesList retrieves a list of tried recipes
func GetTriedRecipesList(c *fiber.Ctx) error {
	return GetRecipesByStatus(c, "tried")
}

func GetNotTriedRecipesList(c *fiber.Ctx) error {
	return GetRecipesByStatus(c, "not-tried")
}
func GetToCookRecipesList(c *fiber.Ctx) error {
	return GetRecipesByStatus(c, "to-cook")
}
func markRecipeStatus(c *fiber.Ctx, status string) error {
	// Get recipeID from params
	recipeID := c.Params("recipeID")

	// Retrieve the recipe from the database
	var recipe models.Recipe
	result := database.DB.Db.First(&recipe, recipeID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Update the recipe status
	recipe.Status = status
	database.DB.Db.Save(&recipe)

	// Return the updated recipe
	return c.JSON(recipe)
}

func GetRecipesByStatus(c *fiber.Ctx, status string) error {
	fmt.Println("Status:", status)

	var recipes []models.Recipe

	// Construct the SQL query dynamically using the status parameter
	sql := "SELECT * FROM recipes WHERE status = '" + status + "' AND deleted_at IS NULL ORDER BY id"

	// Execute the dynamically constructed SQL query
	err := database.DB.Db.Raw(sql).Find(&recipes).Error
	if err != nil {
		fmt.Println("Error fetching recipes by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(recipes)
}

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

// DeleteRecipe deletes a recipe from the database
func DeleteRecipe(c *fiber.Ctx) error {
	// Get recipeID from params
	recipeID := c.Params("recipeID")

	// Retrieve the recipe from the database
	var recipe models.Recipe
	result := database.DB.Db.First(&recipe, recipeID)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	// Delete the recipe from the database
	database.DB.Db.Delete(&recipe)

	// Customize the confirmation message for recipe deletion
	title := "Recipe deleted successfully"
	subtitle := "You can manage your recipe collection"
	return ConfirmationView(c, title, subtitle)
}
func GetToBeCookedRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "to-cook").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching recipes by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(recipes)
}

func GetCookingRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "cooking").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching recipes by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(recipes)
}

func GetTriedRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "tried").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching recipes by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(recipes)
}

func GetNotTriedRecipes(c *fiber.Ctx) error {
	fmt.Println("Status:", "not-tried")

	var recipes []models.Recipe
	sql := "SELECT * FROM recipes WHERE status = ? AND recipes.deleted_at IS NULL ORDER BY id"

	err := database.DB.Db.Raw(sql, "not-tried").Find(&recipes).Error
	if err != nil {
		fmt.Println("Error fetching recipes by status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	return c.JSON(recipes)
}
