package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type MyGoal struct {
	gorm.Model
	UserID      int        `gorm:"not null"`
	Goal        string     `gorm:"not null"`
	GenreID     int        `gorm:"not null"`
	LimitDate   time.Time  `sql:"not null;type:date"`
	IgnoreMe    string     `gorm:"-"`
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

