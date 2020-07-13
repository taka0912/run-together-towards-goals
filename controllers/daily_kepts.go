package controllers

import (
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	ru := models.NewUserRepository()
	user := ru.GetLoginUser(sessions.Default(c).Get("UserId"))
	if user.ID == 0 {
		c.Redirect(http.StatusUnauthorized, "/logout")
	}
	userId := int(user.ID)

	// TODO
	//loginUserId, err := GetLoginUserId(c)
	//if err != nil {
	//	c.Redirect(http.StatusMovedPermanently, "/logout")
	//}
	keep, _    := c.GetPostForm("keep")
	problem, _ := c.GetPostForm("problem")
	try, _     := c.GetPostForm("try")

	// TODO
	r := models.NewDailyKptRepository()
	r.Add(&models.DailyKpt{UserID:userId, Keep:keep, Problem:problem, Try:try})

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
