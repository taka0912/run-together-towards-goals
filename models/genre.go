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
	defer db.Close()
	db.Create(genre)
}

// Edit ...
func (o *Genre) Edit(genre Genre) {
	db := Open()
	defer db.Close()
	genre.UpdatedAt = time.Now()
	db.Save(genre)
}

// GetAll ...
func (o *Genre) GetAll() []Genre {
	var genres []Genre
	db := Open()
	defer db.Close()
	db.Find(&genres)
	return genres
}

// GetOne ...
func (o *Genre) GetOne(id int) Genre {
	var genre Genre
	db := Open()
	defer db.Close()
	db.First(&genre, id)
	return genre
}

// Delete ...
func (o *Genre) Delete(id int) {
	var genre Genre
	db := Open()
	defer db.Close()
	db.First(&genre, id)
	db.Delete(&genre)
}

// GenreMigration ...
func (o *Genre) GenreMigration() {
	db := Open()
	defer db.Close()
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
}
