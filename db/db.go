package db

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

const (
	Driver = "sqlite3"
	DbName  = "my_goal.db"
)

// dbInit...
func Init() *gorm.DB {
	db, err := gorm.Open(Driver, DbName)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	db.AutoMigrate(&models.User{}, &models.DailyKpt{}, &models.MyGoal{}, &models.Genre{})

	defer db.Close()
	return db
}

func Open() *gorm.DB {
	db, err := gorm.Open(Driver, DbName)
	if err != nil {
		panic(err)
	}
	return db
}