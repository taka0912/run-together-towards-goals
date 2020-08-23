package models

import (
	"github.com/jinzhu/gorm"
	"time"
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
	defer db.Close()
	db.Create(goal)
}

// Edit ...
func (o *Goal) Edit(goal Goal) {
	db := Open()
	defer db.Close()
	goal.UpdatedAt = time.Now()
	db.Save(goal)
}

// GetAll ...
func (o *Goal) GetAll() []Goal {
	var goals []Goal
	db := Open()
	defer db.Close()
	db.Find(&goals)
	return goals
}

// GetOne ...
func (o *Goal) GetOne(id int) Goal {
	var goal Goal
	db := Open()
	defer db.Close()
	db.Preload("TodoLists").Find(&goal, id)
	return goal
}

// Delete ...
func (o *Goal) Delete(id int) {
	var goal Goal
	db := Open()
	defer db.Close()
	db.First(&goal, id)
	db.Delete(&goal)
}

// Count ...
func (o *Goal) Count() int {
	db := Open()
	defer db.Close()
	var count = 0
	db.Table("goals").Count(&count)
	return count
}

// GetByUserID ...
func (o *Goal) GetByUserID(id int) []Goal {
	var goals []Goal
	db := Open()
	db.Where("user_id = ?", id).Find(&goals)
	db.Close()
	return goals
}
