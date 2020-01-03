package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jesusantguerrero/goblog/core/models"
)

// Post a type
type Post struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

type CommentModel struct {
	models.Model
}

var updatedFields = []string{"title", "content"}

var schema = `
CREATE TABLE posts IF NOT EXISTS (
	id int autoincrement,
	title varchar(100),
	content text
)
`

// NewModel create new Instance of model
func NewModel() *models.Model {
	model := models.NewModel()
	posts := []Post{}
	model.SetList(&posts)
	return model
}
