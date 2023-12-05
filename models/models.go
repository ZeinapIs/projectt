//models/models.go

package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Status      string `json:"status"` // "read", "to read", "currently reading", "did not finish"
}
