// handlers/recipes.go

package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/ZeinapIs/projectt/database"
	"github.com/ZeinapIs/projectt/models"
	"github.com/gofiber/fiber/v2"
)

// IndexView renders the main page with a list of recipes
func IndexView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Recipe Book",
	})
}

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

// AddNewRecipeView renders the page to add a new recipe
func AddNewRecipeView(c *fiber.Ctx) error {
	return c.Render("add", fiber.Map{
		"Title": "Add New Recipe",
	})
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

	// Set a default value for the Status field
	newRecipe.Status = "default-status"

	// Insert the new recipe into the database
	database.DB.Db.Create(&newRecipe)

	// Redirect to the main page after adding a new recipe
	return c.Redirect("/")
}

// RecipeDetailsView renders the page with details of a specific recipe
func RecipeDetailsView(c *fiber.Ctx) error {
	return c.Render("details", fiber.Map{
		"Title": "Recipe Details",
	})
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

// MarkAsCookedView renders the confirmation page for marking a recipe as cooked
func MarkAsCookedView(c *fiber.Ctx) error {
	return c.Render("mark_cooked", fiber.Map{
		"Title": "Mark Recipe as Cooked",
	})
}

// MarkAsCooked marks a recipe as cooked
func MarkAsCooked(c *fiber.Ctx) error {
	return markRecipeStatus(c, "cooked")
}

// MarkAsFavoriteView renders the confirmation page for marking a recipe as favorite
func MarkAsFavoriteView(c *fiber.Ctx) error {
	return c.Render("mark_favorite", fiber.Map{
		"Title": "Mark Recipe as Favorite",
	})
}

// MarkAsFavorite marks a recipe as a favorite
func MarkAsFavorite(c *fiber.Ctx) error {
	return markRecipeStatus(c, "favorite")
}

// GetCookedRecipesListView renders the page with a list of cooked recipes
func GetCookedRecipesListView(c *fiber.Ctx) error {
	return c.Render("cooked_recipes", fiber.Map{
		"Title": "Cooked Recipes",
	})
}

// GetCookedRecipesList retrieves a list of cooked recipes
func GetCookedRecipesList(c *fiber.Ctx) error {
	return getRecipesByStatus(c, "cooked")
}

// GetFavoriteRecipesListView renders the page with a list of favorite recipes
func GetFavoriteRecipesListView(c *fiber.Ctx) error {
	return c.Render("favorite_recipes", fiber.Map{
		"Title": "Favorite Recipes",
	})
}

// GetFavoriteRecipesList retrieves a list of favorite recipes
func GetFavoriteRecipesList(c *fiber.Ctx) error {
	return getRecipesByStatus(c, "favorite")
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

func getRecipesByStatus(c *fiber.Ctx, status string) error {
	// Implement the logic to fetch recipes by status from the database (similar to getBooksByStatus)
	var recipes []models.Recipe
	database.DB.Db.Where("status = ?", status).Find(&recipes)
	return c.JSON(recipes)
}

// DeleteRecipeView renders the confirmation page for deleting a recipe
func DeleteRecipeView(c *fiber.Ctx) error {
	return c.Render("delete", fiber.Map{
		"Title": "Delete Recipe",
	})
}
func SearchRecipesByTitleView(c *fiber.Ctx) error {
	return c.Render("search_title", fiber.Map{
		"Title": "Search Recipes by Title",
	})
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

// SearchRecipesByIngredientsView renders the page for searching recipes by ingredients
func SearchRecipesByIngredientsView(c *fiber.Ctx) error {
	return c.Render("search_ingredients", fiber.Map{
		"Title": "Search Recipes by Ingredients",
	})
}

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

// SearchRecipesByInstructionsView renders the page for searching recipes by instructions
func SearchRecipesByInstructionsView(c *fiber.Ctx) error {
	return c.Render("search_instructions", fiber.Map{
		"Title": "Search Recipes by Instructions",
	})
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

	// Redirect to the main page after deleting a recipe
	return c.Redirect("/")
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
