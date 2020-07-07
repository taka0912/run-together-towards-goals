package models

import (
	"github.com/jinzhu/gorm"
)

type MyGoal struct {
	gorm.Model
	UserID   int    `gorm:"not null"`
	GenreID  int    `gorm:"not null"`
	Goal     string `gorm:"not null"`
	IgnoreMe string `gorm:"-"`
	Todo Todo `gorm:"foreignkey:ID;association_foreignkey:GoalId"`
}

type MyGoals struct {
	ID     int
	UserID int
	Goal   string
	//Todo   []Todo
	Todo Todo `gorm:"foreignkey:ID;association_foreignkey:GoalId"`
}

// NewMyGoalRepository ...
func NewMyGoalRepository() MyGoal {
	return MyGoal{}
}

// DB追加
func (o *MyGoal) Add(myGoal *MyGoal) {
	db := Open()
	db.Create(myGoal)
	defer db.Close()
}

// DB更新
func (o *MyGoal) Edit(myGoal MyGoal) {
	db := Open()
	db.Save(myGoal)
	db.Close()
}

// DB全取得
func (o *MyGoal) GetAll() []MyGoal {
	db := Open()
	var myGoals []MyGoal
	db.Find(&myGoals)
	db.Close()
	return myGoals
}

// DB一つ取得
func (o *MyGoal) GetOne(id int) MyGoal {
	db := Open()
	var myGoal MyGoal
	db.First(&myGoal, id)
	db.Close()
	return myGoal
}

// GetByUserId
func (o *MyGoal) GetByUserId(userId int) MyGoals {
	db := Open()
	var myGoals MyGoals
	//db.Table("my_goals").
	//	Select("my_goals.*, todos.*").
	//	Joins("left JOIN todos ON my_goals.id = todos.goal_id").
	//	Where("user_id = ?", userId).
	//	Find(&myGoals)
	db.Where(&myGoals, userId).Related(&myGoals.Todo, "Todo")
	db.Close()
	return myGoals
}

// DB削除
func (o *MyGoal) Delete(id int) {
	db := Open()
	var myGoal MyGoal
	db.First(&myGoal, id)
	db.Delete(&myGoal)
	db.Close()
}

// Count...
func (o *MyGoal) Count() int {
	db := Open()
	var count = 0
	db.Table("my_goals").Count(&count)
	db.Close()
	return count
}
