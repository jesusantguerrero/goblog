package models

import (
	"log"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	ID      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title"`
	Content string `json:"content" db:"content"`
}

var updatedFields = []string{"title, content"}

var schema = `
CREATE TABLE posts IF NOT EXISTS (
	id int autoincrement,
	title varchar(100),
	content text
)
`

var db *sqlx.DB

// GetPosts : Get all the posts from database

func Get() []Post {
	db = connect()
	posts := []Post{}
	db.Select(&posts, "select * from posts")
	return posts
}

func Save(post *Post) *Post {
	db = connect()
	sql := "insert into posts(" + strings.Join(updatedFields, ", :") + ") VALUES (" + strings.Join(updatedFields, ", :") + ")"
	db.NamedExec(sql, post)
	return post
}

func GetOne(id int) *Post {
	db = connect()
	post := &Post{}
	err := db.Get(post, "select * from posts where id=?", id)
	if err != nil {
		log.Panic(err)
	}
	return post
}

func update(post *Post, id int) *Post {
	db = connect()
	var toUpdate []string

	for index := 0; index < len(updatedFields); index++ {
		toUpdate = append(toUpdate, updatedFields[index]+"=:"+updatedFields[index])
	}
	idString := strconv.Itoa(id)
	sql := "update  set " + strings.Join(updatedFields, ", ") + " where id=" + idString
	db.NamedExec(sql, post)
	return post
}

func connect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root@(localhost:3307)/test")
	if err != nil {
		log.Panic(err)
	}

	return db
}
