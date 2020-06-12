package api

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

type MyGoal struct {
	gorm.Model
	UserID      string   `json:"user_id"`
	Goal        string   `json:"goal"`
	GenreID     string   `json:"genre_id"`
	LimitDate   string   `json:"limit_date"`
}

// PostDailyKpt...
func (h *Handler) SetMyGoal(c *gin.Context) {
	var myGoal MyGoal
	c.BindJSON(&myGoal)

	userID, _ := strconv.Atoi(myGoal.UserID)
	genreID, _ := strconv.Atoi(myGoal.GenreID)
	limitDate, _ := time.Parse("2006/01/02", myGoal.LimitDate)

	r := models.NewMyGoalRepository()
	r.Add(&models.MyGoal{
		UserID:  userID,
		Goal:    myGoal.Goal,
		GenreID:    genreID,
		LimitDate:    limitDate,
	})

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg" : "Created",
		"id"  : r.Count(),
	})
}

// EditMyGoal...
func (h *Handler) EditMyGoal(c *gin.Context) {
	r := models.NewMyGoalRepository()

	id := c.DefaultQuery("id", "0")
	myGoalId, _ := strconv.Atoi(id)
	newMyGoal := r.GetOne(myGoalId)

	if newMyGoal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code": http.StatusNotFound,
			"msg" : "Not Found",
		})
		return
	}

	var myGoal MyGoal
	c.BindJSON(&myGoal)

	newMyGoal.Goal         = myGoal.Goal
	newMyGoal.GenreID, _   = strconv.Atoi(myGoal.GenreID)
	newMyGoal.LimitDate, _ = time.Parse("2006/01/02", myGoal.LimitDate)

	r.Edit(newMyGoal)

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg" : "Created",
		"id"  : newMyGoal.ID,
	})
}
