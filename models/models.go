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
