// models.go
package models

import (
	"github.com/jinzhu/gorm"
)

// Recipe represents a recipe model
type Recipe struct {
	gorm.Model
	Title        string `json:"title"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
	Status       string `json:"status"`
}

// RecipeResponse represents a simplified recipe response without GORM fields
type RecipeResponse struct {
	Title        string `json:"title"`
	Ingredients  string `json:"ingredients"`
	Instructions string `json:"instructions"`
	Status       string `json:"status"`
}

// CreateRecipeResponse creates a RecipeResponse from a Recipe
func CreateRecipeResponse(recipe Recipe) RecipeResponse {
	return RecipeResponse{
		Title:        recipe.Title,
		Ingredients:  recipe.Ingredients,
		Instructions: recipe.Instructions,
		Status:       recipe.Status,
	}
}
