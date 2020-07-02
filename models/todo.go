package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	GoalId            int        `gorm:"not null"`
	RequiredElements  string     `gorm:"not null"`
	SpecificGoal      string     `gorm:"not null"`
	LimitDate         time.Time  `sql:"not null;type:date"`
	IgnoreMe          string     `gorm:"-"`
}

// NewTodoRepository ...
func NewTodoRepository() Todo {
	return Todo{}
}

// DB追加
func (o *Todo) Add(todo *Todo) {
	db := Open()
	db.Create(todo)
	defer db.Close()
}

// DB更新
func (o *Todo) Edit(todo Todo) {
	db := Open()
	db.Save(todo)
	db.Close()
}

// DB全取得
func (o *Todo) GetAll() []Todo {
	db := Open()
	var todos []Todo
	db.Find(&todos)
	db.Close()
	return todos
}

// DB一つ取得
func (o *Todo) GetOne(id int) Todo {
	db := Open()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}

// DB削除
func (o *Todo) Delete(id int) {
	db := Open()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

// Count...
func (o *Todo) Count() int {
	db := Open()
	var count = 0
	db.Table("todos").Count(&count)
	db.Close()
	return count
}

