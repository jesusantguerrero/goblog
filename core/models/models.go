package models

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Resource interface{}

type Model struct {
	ResourceType  Resource
	resourceName  string
	schema        string
	updatedFields []string
	list          interface{}
	DB            *sqlx.DB
}

// SetList set empty list to fill when calling Get
func (m *Model) SetList(list interface{}) {
	m.list = list
}

// Connect to the database
func (m *Model) Connect() {
	db, err := sqlx.Connect("mysql", "root@(localhost:3306)/test")
	if err != nil {
		log.Panic(err)
	}

	m.DB = db
}

// Get get list of resource from db
func (m Model) Get() interface{} {
	m.Connect()
	fmt.Println(reflect.TypeOf(m.list))
	err := m.DB.Select(m.list, "select * from posts")
	handleError(err)
	m.DB.Close()
	return m.list
}

// Save - seve the list
func (m Model) Save(post interface{}) interface{} {
	m.Connect()
	sql := "insert into posts(" + strings.Join(m.updatedFields, ",") + ") VALUES (:" + strings.Join(m.updatedFields, ", :") + ")"
	result, err := m.DB.NamedExec(sql, post)
	handleError(err)
	lastID, _ := result.LastInsertId()
	m.DB.Close()
	lastIDInt := int(lastID)
	newResource := m.GetOne(lastIDInt)
	return newResource
}

// GetOne - get one row
func (m Model) GetOne(id int) interface{} {
	m.Connect()
	post := &m.ResourceType
	err := m.DB.Get(post, "select * from posts where id=?", id)
	if err != nil {
		log.Panic(err)
	}
	m.DB.Close()
	return post
}

// Update - update a row
func (m Model) Update(post interface{}, id int) interface{} {
	m.Connect()
	var toUpdate []string

	for index := 0; index < len(m.updatedFields); index++ {
		fields := m.updatedFields[index] + "=:" + m.updatedFields[index]
		toUpdate = append(toUpdate, fields)
	}
	idString := strconv.Itoa(id)
	sql := "UPDATE posts SET " + strings.Join(toUpdate, ", ") + " WHERE id=" + idString
	_, err := m.DB.NamedExec(sql, post)
	if err != nil {
		panic(err)
	}
	return post
}

func (m Model) Delete(id int) {
	m.Connect()
	sql := "DELETE FROM posts WHERE id=?"
	m.DB.MustExec(sql, id)
	m.DB.Close()
}

func NewModel() *Model {
	return &Model{}
}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
