package db

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

// dbInit...
func Init() *gorm.DB {
	db := models.Open()

	db.LogMode(true)

	//db.Exec("DROP TABLE genres")
	db.AutoMigrate(
		&models.User{},
		&models.DailyKpt{},
		&models.Goal{},
		&models.Genre{},
		&models.KptReactionHistory{},
		&models.TodoList{},
	)

	r := models.NewGenreRepository()
	r.GenreMigration()

	defer db.Close()
	return db
}
