package models

import (
	BaseModel "github.com/jesusantguerrero/goblog/core/models"
)

// Data - types for post
type Data struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
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
func NewModel() *BaseModel.Model {
	model := BaseModel.NewModel()
	comments := []*Data{}
	comment := new(Data)
	model.Storable = PostModel{}
	model.SetList(&comments)
	model.SetResourceName("posts")
	model.SetUpdateFields(updatedFields)
	model.SetResource(comment)
	return model
}

// PostModel model to manage the posts
type PostModel struct {
}

// GetNewData data from this model data type
func (pm PostModel) GetNewData() interface{} {
	data := &Data{}
	return data
}
