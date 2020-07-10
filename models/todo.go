package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TodoList struct {
	gorm.Model
	GoalID           int       `gorm:"not null"`
	RequiredElements string    `gorm:"not null"`
	SpecificGoal     string    `gorm:"not null"`
	LimitDate        time.Time `sql:"not null;type:date"`
	IgnoreMe         string    `gorm:"-"`
}

// NewTodoListRepository ...
func NewTodoListRepository() TodoList {
	return TodoList{}
}

// DB追加
func (o *TodoList) Add(todoList *TodoList) {
	db := Open()
	db.Create(todoList)
	defer db.Close()
}

// DB更新
func (o *TodoList) Edit(todoList TodoList) {
	db := Open()
	todoList.UpdatedAt = time.Now()
	db.Save(todoList)
	db.Close()
}

// DB全取得
func (o *TodoList) GetAll() []TodoList {
	db := Open()
	var todoLists []TodoList
	db.Find(&todoLists)
	db.Close()
	return todoLists
}

// DB一つ取得
func (o *TodoList) GetOne(id int) TodoList {
	db := Open()
	var todoList TodoList
	db.First(&todoList, id)
	db.Close()
	return todoList
}

// DB削除
func (o *TodoList) Delete(id int) {
	db := Open()
	var todoList TodoList
	db.First(&todoList, id)
	db.Delete(&todoList)
	db.Close()
}

// Count...
func (o *TodoList) Count() int {
	db := Open()
	var count = 0
	db.Table("todoLists").Count(&count)
	db.Close()
	return count
}
