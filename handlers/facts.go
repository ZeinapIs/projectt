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
	var recipe models.Recipe
	recipeID := c.Params("recipeID")

	result := database.DB.Db.Where("id = ?", recipeID).First(&recipe)
	if result.Error != nil {

		return c.Render("error", fiber.Map{
			"Title": "Recipe Not Found",
			"Error": "The requested recipe does not exist or might have been deleted.",
		})
	}

	return c.Render("show", fiber.Map{
		"Title":  "Recipe Details",
		"Recipe": recipe,
	})
}
func EditRecipe(c *fiber.Ctx) error {
	var recipe models.Recipe
	recipeID := c.Params("recipeID")

	result := database.DB.Db.Where("id = ?", recipeID).First(&recipe)
	if result.Error != nil {
		return c.Render("error", fiber.Map{
			"Title": "Recipe Not Found",
			"Error": "The requested recipe does not exist or might have been deleted.",
		})
	}

	return c.Render("edit", fiber.Map{
		"Title":    "Edit Recipe",
		"Subtitle": "Edit your recipe details",
		"Recipe":   recipe,
	})
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

func UpdateRecipe(c *fiber.Ctx) error {
	recipe := new(models.Recipe)
	recipeID := c.Params("recipeID") // Make sure 'recipeID' matches the parameter name in your route

	// Parsing the request body
	if err := c.BodyParser(recipe); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	// Write updated values to the database
	result := database.DB.Db.Model(&recipe).Where("id = ?", recipeID).Updates(recipe)
	if result.Error != nil {
		// If there's an error updating the recipe, redirect to the edit recipe view
		return EditRecipe(c)
	}

	// If the update is successful, redirect to the show recipe view
	return GetRecipeDetails(c)
}

// DeleteRecipe removes a recipe from the database
func DeleteRecipe(c *fiber.Ctx) error {
	recipeID := c.Params("recipeID")

	result := database.DB.Db.Delete(&models.Recipe{}, recipeID)
	if result.Error != nil {
		return NotFound(c) // Ensure you have a NotFound handler defined
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Recipe not found"})
	}

	return GetAllRecipes(c) // Redirect to a function that lists all recipes
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
func GetToCookRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	// Fetch recipes with status "to-cook" from your database
	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "to-cook").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching to-cook recipes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Pass the recipes to the to-cook.html template
	return c.Render("to-cook", fiber.Map{
		"Recipes": recipes,
	})
}
func GetCookingRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	// Fetch recipes with status "cooking" from your database
	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "cooking").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching cooking recipes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Pass the recipes to the cooking.html template
	return c.Render("cooking", fiber.Map{
		"Recipes": recipes,
	})
}

func GetTriedRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	// Fetch recipes with status "tried" from your database
	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "tried").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching tried recipes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Pass the recipes to the tried.html template
	return c.Render("tried", fiber.Map{
		"Recipes": recipes,
	})
}

func GetNotTriedRecipes(c *fiber.Ctx) error {
	var recipes []models.Recipe

	// Fetch recipes with status "not-tried" from your database
	sql := "SELECT * FROM recipes WHERE status = $1 AND deleted_at IS NULL ORDER BY id"
	err := database.DB.Db.Raw(sql, "not-tried").Find(&recipes).Error

	if err != nil {
		fmt.Println("Error fetching not-tried recipes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Pass the recipes to the not-tried.html template
	return c.Render("not-tried", fiber.Map{
		"Recipes": recipes,
	})
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
func GetAllRecipesAsJSON(c *fiber.Ctx) error {
	fmt.Println("Fetching all recipes for JSON")
	var recipes []models.Recipe
	if err := database.DB.Db.Find(&recipes).Error; err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch recipes"})
	}
	return c.JSON(recipes)
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
func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}
