package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
)

type DailyKpt struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Keep    string `json:"keep"`
	Problem string `json:"problem"`
	Try     string `json:"try"`
}

// PostDailyKpt...
func (h *Handler) PostDailyKpt(c *gin.Context) {
	var dailyKpt DailyKpt
	c.BindJSON(&dailyKpt)

	userID, _ := strconv.Atoi(dailyKpt.UserID)

	r := models.NewDailyKptRepository()
	r.Add(&models.DailyKpt{
		UserID:  userID,
		Keep:    dailyKpt.Keep,
		Problem: dailyKpt.Problem,
		Try:     dailyKpt.Try,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Created",
		"id":   r.Count(),
	})
}

// GetDailyKpts...
func (h *Handler) GetDailyKpts(c *gin.Context) {
	r := models.NewDailyKptRepository()

	c.JSON(http.StatusOK, gin.H{
		"daily_kpts": r.GetAll(),
	})
}

// IncreaseGood...
func (h *Handler) IncreaseGood(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	r := models.NewDailyKptRepository()
	dailyKpt := r.GetOne(id)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}

	dailyKpt.Good += 1

	r.Edit(dailyKpt)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   dailyKpt.ID,
	})
}

// IncreaseFight...
func (h *Handler) IncreaseFight(c *gin.Context) {
	id, _ := strconv.Atoi(c.DefaultQuery("id", "0"))
	r := models.NewDailyKptRepository()
	dailyKpt := r.GetOne(id)

	if id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "Bad Request",
		})
		return
	}

	dailyKpt.Fight += 1

	r.Edit(dailyKpt)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "Update",
		"id":   dailyKpt.ID,
	})
}
