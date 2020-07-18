package models

import (
	"github.com/jinzhu/gorm"
)

type Goal struct {
	gorm.Model
	UserID   int    `gorm:"not null"`
	GenreID  int    `gorm:"not null"`
	GoalName string `gorm:"not null"`
	IgnoreMe string `gorm:"-"`
	TodoLists    []TodoList
}

// NewGoalRepository...
func NewGoalRepository() Goal {
	return Goal{}
}

// DB追加
func (o *Goal) Add(goal *Goal) {
	db := Open()
	db.Create(goal)
	defer db.Close()
}

// DB更新
func (o *Goal) Edit(goal Goal) {
	db := Open()
	db.Save(goal)
	db.Close()
}

// DB全取得
func (o *Goal) GetAll() []Goal {
	db := Open()
	var goals []Goal
	db.Find(&goals)
	db.Close()
	return goals
}

// DB一つ取得
func (o *Goal) GetOne(id int) Goal {
	db := Open()
	var goal Goal
	//db.Debug().First(&goal, id).Related(&goal.TodoLists, "GoalID")
	db.Preload("TodoLists").Find(&goal, id)
	//db.Model(&goal)
	//db.Debug().First(&user, id)
	//db.Model(&user).Related(&user.Goals)
	db.Close()
	return goal
}

// DB削除
func (o *Goal) Delete(id int) {
	db := Open()
	var goal Goal
	db.First(&goal, id)
	db.Delete(&goal)
	db.Close()
}

// Count...
func (o *Goal) Count() int {
	db := Open()
	var count = 0
	db.Table("goals").Count(&count)
	db.Close()
	return count
}
