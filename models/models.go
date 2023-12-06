// models/models.go

package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Title        string         `json:"title"`
	Ingredients  string         `json:"ingredients"`
	Instructions string         `json:"instructions"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
