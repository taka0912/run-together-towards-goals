package models

import (
	"github.com/jinzhu/gorm"
)

// Goal ...
type Goal struct {
	gorm.Model
	UserID      int    `gorm:"not null"`
	GenreID     int    `gorm:"not null"`
	GoalName    string `gorm:"not null"`
	DisplayFlag int    `gorm:"not null"`
	IgnoreMe    string `gorm:"-"`
	TodoLists   []TodoList
}

const (
	// DisplayFlagFalse ... 0：非表示
	DisplayFlagFalse = iota
	// DisplayFlagTrue ... 1：表示
	DisplayFlagTrue
)

// NewGoalRepository ...
func NewGoalRepository() Goal {
	return Goal{}
}

// Add ...
func (o *Goal) Add(goal *Goal) {
	db := Open()
	db.Create(goal)
	defer db.Close()
}

// Edit ...
func (o *Goal) Edit(goal Goal) {
	db := Open()
	db.Save(goal)
	db.Close()
}

// GetAll ...
func (o *Goal) GetAll() []Goal {
	db := Open()
	var goals []Goal
	db.Find(&goals)
	db.Close()
	return goals
}

// GetOne ...
func (o *Goal) GetOne(id int) Goal {
	db := Open()
	var goal Goal
	db.Preload("TodoLists").Find(&goal, id)
	db.Close()
	return goal
}

// Delete ...
func (o *Goal) Delete(id int) {
	db := Open()
	var goal Goal
	db.First(&goal, id)
	db.Delete(&goal)
	db.Close()
}

// Count ...
func (o *Goal) Count() int {
	db := Open()
	var count = 0
	db.Table("goals").Count(&count)
	db.Close()
	return count
}

// GetByUserID ...
func (o *Goal) GetByUserID(id int) []Goal {
	db := Open()
	var goals []Goal
	db.Where("user_id = ?", id).Find(&goals)
	db.Close()
	return goals
}
