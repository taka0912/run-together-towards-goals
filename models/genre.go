package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Genre ...
type Genre struct {
	gorm.Model
	GenreName string `gorm:"not null"`
	IgnoreMe  string `gorm:"-"`
}

// NewGenreRepository ...
func NewGenreRepository() Genre {
	return Genre{}
}

// Add ...
func (o *Genre) Add(genre *Genre) {
	db := Open()
	db.Create(genre)
	defer db.Close()
}

// Edit ...
func (o *Genre) Edit(genre Genre) {
	db := Open()
	genre.UpdatedAt = time.Now()
	db.Save(genre)
	db.Close()
}

// GetAll ...
func (o *Genre) GetAll() []Genre {
	db := Open()
	var genres []Genre
	db.Find(&genres)
	db.Close()
	return genres
}

// GetOne ...
func (o *Genre) GetOne(id int) Genre {
	db := Open()
	var genre Genre
	db.First(&genre, id)
	db.Close()
	return genre
}

// Delete ...
func (o *Genre) Delete(id int) {
	db := Open()
	var genre Genre
	db.First(&genre, id)
	db.Delete(&genre)
	db.Close()
}

// GenreMigration ...
func (o *Genre) GenreMigration() {
	db := Open()
	r := NewGenreRepository()

	var count = 0
	db.Table("genres").Count(&count)
	if count == 0 {
		r.Add(&Genre{GenreName: "ダイエット"})
		r.Add(&Genre{GenreName: "筋トレ"})
		r.Add(&Genre{GenreName: "健康"})
		r.Add(&Genre{GenreName: "プログラミング"})
		r.Add(&Genre{GenreName: "資格勉強"})
		r.Add(&Genre{GenreName: "副業"})
		r.Add(&Genre{GenreName: "語学"})
		r.Add(&Genre{GenreName: "その他"})
	}
	db.Close()
}
