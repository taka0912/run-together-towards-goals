package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hariNEzuMI928/run-together-towards-goals/models"
	"net/http"
	"time"
)

// GetGenres...
func (h *Handler) GetAllMonthlyPlans(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	monthlyPlans := r.GetAll()

	c.HTML(http.StatusOK, "monthly_plans.html", gin.H{
		"monthlyPlans": monthlyPlans,
	})
}

// AddGenres...
func (h *Handler) AddMonthlyPlans(c *gin.Context) {
	r := models.NewMonthlyPlanRepository()
	loginUserId, err := GetLoginUserId(c)
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, "/logout")
	}
	r.UserID = loginUserId
	r.GoalID = 0
	month, _ := c.GetPostForm("Month")
	r.Month, _ = time.Parse("2006-01", month)
	r.KeepInLastMonth, _ = c.GetPostForm("KeepInLastMonth")
	r.ProblemInLastMonth, _ = c.GetPostForm("ProblemInLastMonth")
	r.GoalAfterHalfYear, _ = c.GetPostForm("GoalAfterHalfYear")
	r.GoalInThisMonth, _ = c.GetPostForm("GoalInThisMonth")
	r.CurrentState, _ = c.GetPostForm("CurrentState")
	r.DailyTodo, _ = c.GetPostForm("DailyTodo")

	r.Add(&r)

	c.Redirect(http.StatusMovedPermanently, "/_monthly_plans")
}

//// GetGenres...
//func (h *Handler) GetGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	genre := r.GetOne(id)
//	c.HTML(http.StatusOK, "genre_edit.html", gin.H{
//		"genre": genre,
//	})
//}
//
//// EditGenres...
//func (h *Handler) EditGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	genre := r.GetOne(id)
//	genre.GenreName, _ = c.GetPostForm("genreName")
//	r.Edit(genre)
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}
//
//// DeleteGenres...
//func (h *Handler) DeleteGenre(c *gin.Context) {
//	r := models.NewGenreRepository()
//	id, _ := strconv.Atoi(c.Param("id"))
//	r.Delete(id)
//	c.Redirect(http.StatusMovedPermanently, "/_genres")
//}
