package models

import (
	BaseModel "github.com/jesusantguerrero/goblog/core/models"
)

// Data - Types of the vaina
type Data struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	PostID   int    `json:"post_id" db:"post_id"`
	Content  string `json:"content" db:"content"`
	EditMode bool   `json:"edit_mode" db:"edit_mode"`
}

var updatedFields = []string{"user_id", "post_id", "content", "edit_mode"}

// NewModel create new Instance of model
func NewModel() *BaseModel.Model {
	model := BaseModel.NewModel()
	comments := []*Data{}
	comment := new(Data)
	model.Storable = CommentModel{}
	model.SetList(&comments)
	model.SetResourceName("comments")
	model.SetUpdateFields(updatedFields)
	model.SetResource(comment)
	return model
}

type CommentModel struct {
}

func (cm CommentModel) GetNewData() interface{} {
	data := &Data{}
	return data
}
