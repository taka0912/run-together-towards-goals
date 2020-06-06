package controllers

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUsers ...
func (h *Handler) GetAllDailyKpts(c *gin.Context) {
	r := models.NewDailyKptRepository()
	dailyKpts := r.GetAll()

	c.HTML(http.StatusOK, "daily_kpts.html", gin.H{
		"dailyKpts": dailyKpts,
	})
}

// DeleteUsers ...
func (h *Handler) DeleteDailyKpt(c *gin.Context) {
	r := models.NewDailyKptRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/daily_kpts")
}

