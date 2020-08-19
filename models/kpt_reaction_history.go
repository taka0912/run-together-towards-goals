package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// KptReactionHistory ...
type KptReactionHistory struct {
	gorm.Model
	KptID    int    `gorm:"not null"`
	UserID   int    `gorm:"not null"`
	Reaction int    `gorm:"not null"`
	IgnoreMe string `gorm:"-"`
}

const (
	// ReactionOthers ...
	ReactionOthers = iota
	// ReactionGood ...
	ReactionGood
	// ReactionFight ...
	ReactionFight
)

// NewKptReactionHistoryRepository ...
func NewKptReactionHistoryRepository() KptReactionHistory {
	return KptReactionHistory{}
}

// Add ...
func (o *KptReactionHistory) Add(kptReactionHistory *KptReactionHistory) {
	db := Open()
	db.Create(kptReactionHistory)
	defer db.Close()
}

// AddReaction ...
func (o *KptReactionHistory) AddReaction(kptID int, userID int, reaction int) {
	db := Open()
	newReaction := NewKptReactionHistoryRepository()
	newReaction.KptID = kptID
	newReaction.UserID = userID
	newReaction.Reaction = reaction

	db.Create(&newReaction)
	defer db.Close()
}

// Edit ...
func (o *KptReactionHistory) Edit(kptReactionHistory KptReactionHistory) {
	db := Open()
	db.Save(kptReactionHistory)
	kptReactionHistory.UpdatedAt = time.Now()
	db.Close()
}

// GetAll ...
func (o *KptReactionHistory) GetAll() []KptReactionHistory {
	db := Open()
	var kptReactionHistorys []KptReactionHistory
	db.Find(&kptReactionHistorys)
	db.Close()
	return kptReactionHistorys
}

// GetOne ...
func (o *KptReactionHistory) GetOne(id int) KptReactionHistory {
	db := Open()
	var kptReactionHistory KptReactionHistory
	db.First(&kptReactionHistory, id)
	db.Close()
	return kptReactionHistory
}

// Delete ...
func (o *KptReactionHistory) Delete(id int) {
	db := Open()
	var kptReactionHistory KptReactionHistory
	db.First(&kptReactionHistory, id)
	db.Delete(&kptReactionHistory)
	db.Close()
}
