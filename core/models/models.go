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
	resource      interface{}
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

// SetResource set resource instance fill when calling Get
func (m *Model) SetResource(resource interface{}) {
	m.resource = resource
}

// SetResourceName set empty list to fill when calling Get
func (m *Model) SetResourceName(name string) {
	m.resourceName = name
}

// SetUpdateFields set empty list to fill when calling Get
func (m *Model) SetUpdateFields(fields []string) {
	m.updatedFields = fields
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
	err := m.DB.Select(m.list, "select * from "+m.resourceName)
	handleError(err)
	m.DB.Close()
	return m.list
}

// Save - seve the list
func (m Model) Save(post interface{}) interface{} {
	m.Connect()
	sql := "insert into " + m.resourceName + "(" + strings.Join(m.updatedFields, ",") + ") VALUES (:" + strings.Join(m.updatedFields, ", :") + ")"
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
	resource := m.resource
	err := m.DB.Get(resource, "select * from "+m.resourceName+" where id=?", id)
	if err != nil {
		log.Panic(err)
	}
	m.DB.Close()
	return resource
}

// Update - update a row
func (m Model) Update(resource interface{}, id int) interface{} {
	m.Connect()
	var toUpdate []string

	for index := 0; index < len(m.updatedFields); index++ {
		fields := m.updatedFields[index] + "=:" + m.updatedFields[index]
		toUpdate = append(toUpdate, fields)
	}
	idString := strconv.Itoa(id)
	sql := "UPDATE " + m.resourceName + " SET " + strings.Join(toUpdate, ", ") + " WHERE id=" + idString
	_, err := m.DB.NamedExec(sql, resource)
	if err != nil {
		panic(err)
	}
	return resource
}

// Delete resource
func (m Model) Delete(id int) {
	m.Connect()
	sql := "DELETE FROM " + m.resourceName + " WHERE id=?"
	m.DB.MustExec(sql, id)
	m.DB.Close()
}

// NewModel create a new model instance
func NewModel() *Model {
	return &Model{}
}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
