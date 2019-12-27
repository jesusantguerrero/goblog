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

var updatedFields = []string{"title", "content"}

var schema = `
CREATE TABLE posts IF NOT EXISTS (
	id int autoincrement,
	title varchar(100),
	content text
)
`

var db *sqlx.DB

// GetPosts : Get all the posts from database

func connect() *sqlx.DB {
	db, err := sqlx.Connect("mysql", "root@(localhost:3307)/test")
	if err != nil {
		log.Panic(err)
	}

	return db
}

func Get() []Post {
	db = connect()
	posts := []Post{}
	db.Select(&posts, "select * from posts")
	db.Close()
	return posts
}

func Save(post *Post) *Post {
	db = connect()
	sql := "insert into posts(" + strings.Join(updatedFields, ", :") + ") VALUES (" + strings.Join(updatedFields, ", :") + ")"
	db.NamedExec(sql, post)
	db.Close()
	return post
}

func GetOne(id int) *Post {
	db = connect()
	post := &Post{}
	err := db.Get(post, "select * from posts where id=?", id)
	if err != nil {
		log.Panic(err)
	}
	db.Close()
	return post
}

func Update(post *Post, id int) *Post {
	db = connect()
	var toUpdate []string

	for index := 0; index < len(updatedFields); index++ {
		fields := updatedFields[index] + "=:" + updatedFields[index]
		toUpdate = append(toUpdate, fields)
	}
	idString := strconv.Itoa(id)
	sql := "UPDATE posts SET " + strings.Join(toUpdate, ", ") + " WHERE id=" + idString
	_, err := db.NamedExec(sql, post)
	if err != nil {
		panic(err)
	}
	return post
}

func Delete(id int) {
	db = connect()
	sql := "DELETE FROM posts WHERE id=?"
	db.MustExec(sql, id)
	db.Close()
}
