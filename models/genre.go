package models

import (
	"github.com/daisuzuki829/run_together_towards_goals/db"
	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	GenreName    string  `gorm:"not null"`
	IgnoreMe     string  `gorm:"-"`
}

// GenreRepository is
type GenreRepository struct {
}

// NewGenreRepository ...
func NewGenreRepository() GenreRepository {
	return GenreRepository{}
}

//DB追加
func (r *GenreRepository) Add(genre *Genre) {
	db := db.Open()
	db.Create(genre)
	defer db.Close()
}

//DB更新
func (r *GenreRepository) Edit(genre Genre) {
	db := db.Open()
	db.Save(genre)
	db.Close()
}

//DB全取得
func (r *GenreRepository) GetAll() []Genre {
	db := db.Open()
	var genres []Genre
	db.Find(&genres)
	db.Close()
	return genres
}

//DB一つ取得
func (r *GenreRepository) GetOne(id int) Genre {
	db := db.Open()
	var genre Genre
	db.First(&genre, id)
	db.Close()
	return genre
}

//DB削除
func (r *GenreRepository) Delete(id int) {
	db := db.Open()
	var genre Genre
	db.First(&genre, id)
	db.Delete(&genre)
	db.Close()
}

