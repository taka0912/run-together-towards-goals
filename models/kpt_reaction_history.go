package models

import (
	"github.com/jinzhu/gorm"
)

type KptReactionHistory struct {
	gorm.Model
	KptId      int     `gorm:"not null"`
	UserId     int     `gorm:"not null"`
	Reaction   int     `gorm:"not null"`
	IgnoreMe   string  `gorm:"-"`
}

const (
	ReactionOthers = iota
	ReactionGood
	ReactionFight
)

// NewKptReactionHistoryRepository...
func NewKptReactionHistoryRepository() KptReactionHistory {
	return KptReactionHistory{}
}

// DB追加
func (o *KptReactionHistory) Add(kptReactionHistory *KptReactionHistory) {
	db := Open()
	db.Create(kptReactionHistory)
	defer db.Close()
}

// DB追加
func (o *KptReactionHistory) AddReaction(kptId int, userId int, reaction int) {
	db := Open()
	newReaction := NewKptReactionHistoryRepository()
	newReaction.KptId    = kptId
	newReaction.UserId   = userId
	newReaction.Reaction = reaction

	db.Create(&newReaction)
	defer db.Close()
}

// DB更新
func (o *KptReactionHistory) Edit(kptReactionHistory KptReactionHistory) {
	db := Open()
	db.Save(kptReactionHistory)
	db.Close()
}

// DB全取得
func (o *KptReactionHistory) GetAll() []KptReactionHistory {
	db := Open()
	var kptReactionHistorys []KptReactionHistory
	db.Find(&kptReactionHistorys)
	db.Close()
	return kptReactionHistorys
}

// DB一つ取得
func (o *KptReactionHistory) GetOne(id int) KptReactionHistory {
	db := Open()
	var kptReactionHistory KptReactionHistory
	db.First(&kptReactionHistory, id)
	db.Close()
	return kptReactionHistory
}

// DB削除
func (o *KptReactionHistory) Delete(id int) {
	db := Open()
	var kptReactionHistory KptReactionHistory
	db.First(&kptReactionHistory, id)
	db.Delete(&kptReactionHistory)
	db.Close()
}


