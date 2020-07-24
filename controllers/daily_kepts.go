package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

// GetAllDailyKpts...
func (h *Handler) GetAllDailyKpts(c *gin.Context) {
	r := models.NewDailyKptRepository()
	dailyKpts := r.GetAll()

	c.HTML(http.StatusOK, "daily_kpts.html", gin.H{
		"dailyKpts": dailyKpts,
	})
}

// AddDailyKpt...
func (h *Handler) AddDailyKpt(c *gin.Context) {
	loginUserId, err := GetLoginUserId(c)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/logout")
	}

	r := models.NewDailyKptRepository()
	r.UserID = loginUserId
	r.Keep, _ = c.GetPostForm("keep")
	r.Problem, _ = c.GetPostForm("problem")
	r.Try, _ = c.GetPostForm("try")

	r.Add(&r)
	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// IncreaseGood...
func (h *Handler) IncreaseGood(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	dailyKpt := r.GetOne(id)
	dailyKpt.Good += 1
	r.Edit(dailyKpt)

	loginUserId, err := GetLoginUserId(c)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/logout")
	}

	rah := models.NewKptReactionHistoryRepository()
	rah.AddReaction(int(dailyKpt.ID), loginUserId, models.ReactionGood)

	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// IncreaseFight...
func (h *Handler) IncreaseFight(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	dailyKpt := r.GetOne(id)
	dailyKpt.Fight += 1
	r.Edit(dailyKpt)

	loginUserId, err := GetLoginUserId(c)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/logout")
	}

	rah := models.NewKptReactionHistoryRepository()
	rah.AddReaction(int(dailyKpt.ID), loginUserId, models.ReactionFight)

	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}

// DeleteDailyKpt...
func (h *Handler) DeleteDailyKpt(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/_daily_kpts")
}
