package controllers

import (
	"github.com/daisuzuki829/run_together_towards_goals/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetGenres ...
func (h *Handler) GetAllGenres(c *gin.Context) {
	r := models.NewGenreRepository()
	genres := r.GetAll()

	c.HTML(http.StatusOK, "genres.html", gin.H{
		"genres": genres,
	})
}

// AddGenres ...
func (h *Handler) AddGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	genreName, _ := c.GetPostForm("genreName")

	r.Add(&models.Genre{GenreName: genreName})

	c.Redirect(http.StatusMovedPermanently, "/genres")
}

// GetGenres ...
func (h *Handler) GetGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	genre := r.GetOne(id)
	c.HTML(http.StatusOK, "genre_edit.html", gin.H{
		"genre": genre,
	})
}

// EditGenres ...
func (h *Handler) EditGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	genre := r.GetOne(id)
	genre.GenreName, _ = c.GetPostForm("genreName")
	r.Edit(genre)
	c.Redirect(http.StatusMovedPermanently, "/genres")
}

// DeleteGenres ...
func (h *Handler) DeleteGenre(c *gin.Context) {
	r := models.NewGenreRepository()
	id, _ := strconv.Atoi(c.Param("id"))
	r.Delete(id)
	c.Redirect(http.StatusMovedPermanently, "/genres")
}

