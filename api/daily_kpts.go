package api

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DailyKpt struct {
	UserID    string  `json:"user_id"`
	Keep      string  `json:"keep"`
	Problem   string  `json:"problem"`
	Try       string  `json:"try"`
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
		"msg" : "Created",
		"id"  : r.Count(),
	})
}
