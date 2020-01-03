package models

import (
	"github.com/jesusantguerrero/goblog/core/models"
)

// Comment a type
type Comment struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	PostID   int    `json:"post_id" db:"post_id"`
	Content  string `json:"content" db:"content"`
	EditMode bool   `json:"edit_mode" db:"edit_mode"`
}

var updatedFields = []string{"user_id", "post_id", "content", "edit_mode"}

// NewModel create new Instance of model
func NewModel() *models.Model {
	model := models.NewModel()
	comments := []Comment{}
	model.SetList(&comments)
	model.SetResourceName("comments")
	model.SetUpdateFields(updatedFields)
	model.SetResource(&Comment{})
	return model
}
