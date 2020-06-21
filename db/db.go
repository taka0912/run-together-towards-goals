package db

import (
	"github.com/daisuzuki829/run-together-towards-goals/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// dbInit...
func Init() *gorm.DB {
	db := models.Open()

	db.LogMode(true)

	db.Exec("DROP TABLE genres")
	db.AutoMigrate(&models.User{}, &models.DailyKpt{}, &models.MyGoal{}, &models.Genre{})

	r := models.NewGenreRepository()
	r.GenreMigration()

	defer db.Close()
	return db
}
