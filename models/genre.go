package models

import (
	"github.com/jinzhu/gorm"
)

type Genre struct {
	gorm.Model
	GenreName    string  `gorm:"not null"`
	IgnoreMe     string  `gorm:"-"`
}

// NewGenreRepository ...
func NewGenreRepository() Genre {
	return Genre{}
}

// DB追加
func (o *Genre) Add(genre *Genre) {
	db := Open()
	db.Create(genre)
	defer db.Close()
}

// DB更新
func (o *Genre) Edit(genre Genre) {
	db := Open()
	db.Save(genre)
	db.Close()
}

// DB全取得
func (o *Genre) GetAll() []Genre {
	db := Open()
	var genres []Genre
	db.Find(&genres)
	db.Close()
	return genres
}

// DB一つ取得
func (o *Genre) GetOne(id int) Genre {
	db := Open()
	var genre Genre
	db.First(&genre, id)
	db.Close()
	return genre
}

// DB削除
func (o *Genre) Delete(id int) {
	db := Open()
	var genre Genre
	db.First(&genre, id)
	db.Delete(&genre)
	db.Close()
}

// Migration
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
