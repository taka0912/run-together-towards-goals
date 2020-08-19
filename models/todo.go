package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TodoList ...
type TodoList struct {
	gorm.Model
	GoalID           int       `gorm:"not null"`
	RequiredElements string    `gorm:"not null"`
	Todo             string    `gorm:"not null"`
	SpecificGoal     string    `gorm:"not null"`
	LimitDate        time.Time `sql:"not null;type:date"`
	IgnoreMe         string    `gorm:"-"`
}

// NewTodoListRepository ...
func NewTodoListRepository() TodoList {
	return TodoList{}
}

// Add ...
func (o *TodoList) Add(todoList *TodoList) {
	db := Open()
	db.Create(todoList)
	defer db.Close()
}

// Edit ...
func (o *TodoList) Edit(todoList TodoList) {
	db := Open()
	todoList.UpdatedAt = time.Now()
	db.Save(todoList)
	db.Close()
}

// GetAll ...
func (o *TodoList) GetAll() []TodoList {
	db := Open()
	var todoLists []TodoList
	db.Find(&todoLists)
	db.Close()
	return todoLists
}

// GetOne ...
func (o *TodoList) GetOne(id int) TodoList {
	db := Open()
	var todoList TodoList
	db.First(&todoList, id)
	db.Close()
	return todoList
}

// Delete ...
func (o *TodoList) Delete(id int) {
	db := Open()
	var todoList TodoList
	db.First(&todoList, id)
	db.Delete(&todoList)
	db.Close()
}

// Count ...
func (o *TodoList) Count() int {
	db := Open()
	var count = 0
	db.Table("todoLists").Count(&count)
	db.Close()
	return count
}
